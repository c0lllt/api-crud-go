package controllers

import (
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
