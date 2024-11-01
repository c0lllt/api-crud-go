package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import para conectar ao MySQL
)

var BancodeDados *sql.DB // Variável global para a conexão com o banco de dados

func Conectar() {

	// Conexão com banco de dados ...
	var err error
	BancodeDados, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/loja_vendas") // Abre conexao com banco
	if err != nil {
		log.Fatal(err)
	}

	err = BancodeDados.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexao bem sucedida!")
}
