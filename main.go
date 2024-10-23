package main

import (
	// pacote para conexao com bando dedaos.

	"loja-vendas/database"
	"loja-vendas/rotas"

	// modelo inicialmente criado para clientes.

	_ "github.com/go-sql-driver/mysql" // para conectar ao MySQL
)

func main() {

	// Conectar o banco
	database.Conectar()

	// Iniciar o servidor na porta 8081
	router := rotas.ConfigurarRotas()
	router.Run(":8081")
}
