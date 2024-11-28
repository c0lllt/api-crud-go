package rotas

import (
	"loja-vendas/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigurarRotas() *gin.Engine {
	router := gin.Default() // cria um novo router..

	//Configurar CORS para permitir as requisições do front
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rotas para CRUD de clientes

	router.POST("/clientes", controllers.CriarCliente)
	router.GET("/clientes", controllers.BuscarClientes)
	router.GET("/clientes/:id", controllers.BuscarClientes)
	router.PUT("/clientes/:id", controllers.AtualizarCliente)
	router.DELETE("/clientes/:id", controllers.DeletarCliente)
	//rota de login
	router.POST("/login", controllers.LogarUsuario)

	// rota para CRUD de Produtos
	router.POST("/produtos", controllers.CriarProduto)
	router.GET("/produtos", controllers.BuscarProduto) //rota para envio das pizzas para o front

	router.GET("/produtos/:nome", controllers.BuscarProduto)
	router.PUT("/produtos/:id", controllers.AtualizarProduto)
	router.DELETE("/produtos/:id", controllers.DeletarProduto)

	return router

}
