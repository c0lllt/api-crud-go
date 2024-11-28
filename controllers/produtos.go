package controllers

import (
	"database/sql"
	"fmt"
	"loja-vendas/database"
	"loja-vendas/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Funçoes para CRUD de produtos

func CriarProduto(c *gin.Context) {
	var novoProduto models.Produtos

	if err := c.ShouldBindBodyWithJSON(&novoProduto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON invalido", "detalhes": err.Error()})
	}

	//validar campos para certificar que nao estejam vazios
	if novoProduto.Nome == "" || novoProduto.Descricao == "" || novoProduto.Valor == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Nome, descrição e valor  são necessarios"})
		return
	}

	//inserir no banco de dados

	insertQuery := "INSERT INTO produto(nome, descricao,valor) VALUES (?,?,?)"
	resultado, err := database.BancodeDados.Exec(insertQuery, novoProduto.Nome, novoProduto.Descricao, novoProduto.Valor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao inserir no banco"})
		return
	}

	// pegar id gerado pelo banco

	id, err := resultado.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao recuperar ID do produto"})
		return
	}

	// retornar dados do cliente criado
	novoProduto.ID = int(id)
	c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado com sucesso!", "Produto": novoProduto})
}

// Func para listrar Produto(s)

func BuscarProduto(c *gin.Context) {
	nome := c.Param("nome")

	var buscar string
	var rows *sql.Rows
	var err error

	//Buscar Produto Especifico
	if nome != "" {
		buscar = "SELECT id, nome, descricao, valor FROM produto WHERE nome =? "
		rows, err = database.BancodeDados.Query(buscar, nome)
		// Buscar todos
	} else {
		buscar = "SELECT id, nome, descricao, valor FROM produto"
		rows, err = database.BancodeDados.Query(buscar)
	}

	if err != nil {
		fmt.Printf("\n E erro: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar produto"})
		return
	}

	defer rows.Close()

	var produtos []models.Produtos
	for rows.Next() {
		var produto models.Produtos
		if err := rows.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Valor); err != nil {
			fmt.Printf("\n Erro: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao ler dados do produto"})
			return
		}
		produtos = append(produtos, produto)
	}

	if len(produtos) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Nenhum produto encontrado"})
		return
	}

	c.JSON(http.StatusOK, produtos)

}

func AtualizarProduto(c *gin.Context) {
	id := c.Param("id")

	var atualizarProd models.Produtos
	if err := c.ShouldBindBodyWithJSON(&atualizarProd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Problema ao Bindar JSON", "detalhes": err.Error()})
		return
	}

	//Validar os campos obrigatorios
	if atualizarProd.Nome == "" || atualizarProd.Descricao == "" || atualizarProd.Valor == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Nome, Descrição e Valores são obrigatorios"})
		return
	}

	//Atualizar produto no banco
	updateQuery := "UPDATE produto SET nome = ?, descricao = ?, valor = ? WHERE id = ?"
	resultado, err := database.BancodeDados.Exec(updateQuery, atualizarProd.Nome, atualizarProd.Descricao, atualizarProd.Valor, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao tentar atualizar o produto."})
		return
	}

	//Verificar se o produto foi atualizado
	rowsAffected, err := resultado.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"erro": "produto nao encontrado"})
		return
	}

	//Json mostrando que deu certo.
	c.JSON(http.StatusOK, gin.H{"Tudo certo": "Produto atualizado com sucesso!"})
}

func DeletarProduto(c *gin.Context) {
	id := c.Param("id")

	deleteQuery := "DELETE FROM produto WHERE id = ?"
	resultado, err := database.BancodeDados.Exec(deleteQuery, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Erro ao tentar deletar produto.", "Motivo": err.Error()})
		return
	}
	//Verifica se foi deletado
	rowsAffected, err := resultado.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Produto nao encontrado ou ja deletado."})
		return
	}

	//json caso deu tudo certo
	c.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso!"})

}
