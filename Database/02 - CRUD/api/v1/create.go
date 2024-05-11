package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/blkzy/Curso-Go-Projetos/Database/2-CRUD/database"
)

// CreateHTTP :: Está função é chamada quando há uma request no "/api/v1/user" (Método POST).
func CreateHTTP(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println(err)
		response.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	idUser, err := database.Create(db, response, body)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte(fmt.Sprintf("Usuário criado com sucesso! ID: %d", idUser)))
}
