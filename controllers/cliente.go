package controllers

import (
	"database/sql"
	"fmt"
	"loja-vendas/database"
	"loja-vendas/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Aqui vao ficar as funções do CRUD cliente

// Função para criar um cliente
func CriarCliente(c *gin.Context) {
	var novoCliente models.Cliente

	// Validação do JSON recebido
	if err := c.ShouldBindJSON(&novoCliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido", "detalhes": err.Error()})
		return
	}

	// Validando os campos manualmente (garantir que não estão vazios)
	if novoCliente.Nome == "" || novoCliente.Cpf == "" || novoCliente.Senha == "" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Nome, Cpf, Senha  são obrigatórios"})
		return
	}

	// Inserir novo cliente no banco de dados
	insertQuery := "INSERT INTO clientes (nome, cpf, senha) VALUES (?, ?, ?)"
	res, err := database.BancodeDados.Exec(insertQuery, novoCliente.Nome, novoCliente.Cpf, novoCliente.Senha)
	if err != nil {
		fmt.Printf("\n E erro: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao inserir no banco"})
		return
	}

	// Recuperar o ID gerado pelo banco
	id, err := res.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao recuperar ID do cliente"})
		return
	}

	// Retornar os dados do cliente criado, incluindo o ID gerado
	novoCliente.ID = int(id)
	c.JSON(http.StatusOK, gin.H{"message": "Cliente adicionado com sucesso", "cliente": novoCliente})
}

// função para listar clientes
func BuscarClientes(c *gin.Context) {
	id := c.Param("id") // Pega o ID da URL (ex: /clientes/1)

	var query string
	var rows *sql.Rows
	var err error

	if id != "" {
		// Busca por cliente específico
		query = "SELECT id_cliente, nome, cpf FROM clientes WHERE id_cliente=?"
		rows, err = database.BancodeDados.Query(query, id)
	} else {
		// Busca todos os clientes
		query = "SELECT id_cliente, nome, cpf FROM clientes"
		rows, err = database.BancodeDados.Query(query)
	}

	if err != nil {
		fmt.Printf("\n E erro: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar clientes"})
		return
	}
	defer rows.Close()

	var clientes []models.Cliente
	for rows.Next() {
		var cliente models.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.Cpf); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao ler dados do cliente"})
			return
		}
		clientes = append(clientes, cliente)
	}

	if len(clientes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum cliente encontrado"})
		return
	}

	c.JSON(http.StatusOK, clientes)
}

func AtualizarCliente(c *gin.Context) {
	id := c.Param("id") // Pegando o ID da URL corretamente

	var clienteAtualizado models.Cliente
	if err := c.ShouldBindJSON(&clienteAtualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido", "detalhes": err.Error()})
		return
	}

	// Validando os campos obrigatórios
	if clienteAtualizado.Nome == "" || clienteAtualizado.Cpf == "" || clienteAtualizado.Senha == "" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Nome, CPF e senha são obrigatórios"})
		return
	}

	// Atualizar cliente no banco de dados
	updateQuery := "UPDATE clientes SET nome = ?, cpf = ?, senha = ? WHERE id_cliente = ?"
	res, err := database.BancodeDados.Exec(updateQuery, clienteAtualizado.Nome, clienteAtualizado.Cpf, clienteAtualizado.Senha, id)
	if err != nil {
		fmt.Printf("\n E erro: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao atualizar cliente"})
		return
	}

	// Verificando se o cliente foi atualizado
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Cliente não encontrado ou sem alterações"})
		return
	}
	//Mandar um json mostrando que o cliente foi atualziado com sucesso
	c.JSON(http.StatusOK, gin.H{"message": "Cliente atualizado com sucesso"})
}

// Função para deletar cliente
func DeletarCliente(c *gin.Context) {
	id := c.Param("id") // Pegando o ID da URL (ex: /clientes/:id)

	// Query para deletar o cliente pelo ID
	deleteQuery := "DELETE FROM clientes WHERE id_cliente = ?"
	res, err := database.BancodeDados.Exec(deleteQuery, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao deletar cliente"})
		return
	}

	// Verificando se o cliente foi deletado
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Cliente não encontrado ou já deletado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente deletado com sucesso"})
}

// função para Logar o usuario

func LogarUsuario(c *gin.Context) {
	var loginData struct {
		Nome  string `json:"nome"`
		Senha string `json:"senha"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	fmt.Println("Dados recebidos:", loginData)

	var clientes models.Cliente
	err := database.BancodeDados.QueryRow("SELECT id_cliente, nome, senha FROM clientes WHERE nome = ? AND senha =?", loginData.Nome, loginData.Senha).Scan(&clientes.ID, &clientes.Nome, &clientes.Senha)
	if err != nil {
		fmt.Printf("\n E erro: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciais inválidas"})
		return
	}

	// Login bem-sucedido
	c.JSON(http.StatusOK, gin.H{"message": "Login realizado com sucesso"})
}
