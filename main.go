package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// framework para criação da api
)

func main() {

	// Iniciar gin gonic
	router := gin.Default() //cria um novo router com middleware padrão, como logger e recuperação de erros. O router é o que controla as rotas da API.

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/loja_vendas") //Este comando abre uma conexão com o banco de dados MySQL.
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close() //Isso garante que a conexão com o banco de dados será fechada assim que a função main terminar, liberando recursos.

	err = db.Ping() // Aqui verificamos se a conexão foi bem-sucedida. Caso haja algum erro, ele será mostrado e o programa será finalizado.
	if err != nil {
		log.Fatal(err)
	}

	//Rota Para Crud de clientes

	router.POST("/clientes", func(criarCliente *gin.Context) {
		criarCliente.JSON(200, gin.H{"message": "Criando Cliente"})
	})

	router.GET("/clientes", func(listarclientes *gin.Context) {
		listarclientes.JSON(200, gin.H{"message": "Listando todos os clientes"})
	})

}
