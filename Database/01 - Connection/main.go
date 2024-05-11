package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
Para se conectar ao mysql é necessário baixar a lib "github.com/go-sql-driver/mysql", (\n)
por mais que a golang por padrão já tenha a opção de sql... (\n)
banco de dados e etc... ela não vem com o driver por padrão.
Quando importamos o driver é necessário por um _(underline) antes. (\n)
Pôs esse aquivo em sí não utilizará/instanciará a biblioteca.
E exatamente por esse motivo que não instanciase o mysql em sí no "sql.Open()", (\n)
nós informamos uma string com o nome do driver sql.

A URI do mysql deve ser assim: "usuário:senha@/database".

*/

func main() {
	database, err := sql.Open("mysql", "root:toortoor@/golang?charset=utf8")

	if err != nil {
		log.Fatal("Ocorreu um erro:", err)
	}

	defer database.Close()

	if err = database.Ping(); err != nil {
		log.Fatal("Não consegui me conectar ao banco de dados!", err)
	}

	fmt.Println("Me conectei ao banco de dados!")

	sql, err := database.Query("select * from users")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sql)

	defer sql.Close()
}
