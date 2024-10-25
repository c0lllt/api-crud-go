package controllers

import (
	"database/sql"
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
	if novoProduto.Nome == "" || novoProduto.Valor == 0 || novoProduto.Quantidade == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Nome, valor e quantidade são necessarios"})
		return
	}

	//inserir no banco de dados

	insertQuery := "INSERT INTO produto(nome,valor,quantidade) VALUES (?,?,?)"
	resultado, err := database.BancodeDados.Exec(insertQuery, novoProduto.Nome, novoProduto.Valor, novoProduto.Quantidade)
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
	id := c.Param("id")

	var buscar string
	var rows *sql.Rows
	var err error

	//Buscar Produto Especifico
	if id != "" {
		buscar = "SELECT id, nome FROM produto WHERE id = ? "
		rows, err = database.BancodeDados.Query(buscar, id)
		// Buscar todos
	} else {
		buscar = "SELECT id, nome, valor, quantidade FROM produto"
		rows, err = database.BancodeDados.Query(buscar)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar produto"})
		return
	}

	defer rows.Close()

	var produtos []models.Produtos
	for rows.Next() {
		var produto models.Produtos
		if err := rows.Scan(&produto.ID, &produto.Nome, &produto.Valor, &produto.Quantidade); err != nil {
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
