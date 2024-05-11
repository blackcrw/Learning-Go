package api

import (
	"net/http"
	"strconv"

	"github.com/blkzy/Curso-Go-Projetos/Database/2-CRUD/database"
	"github.com/gorilla/mux"
)

// SearchHTTP :: Está função será chamada quando haver uma request com id "/api/v1/user/{id}". (Método GET).
func SearchHTTP(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request) // Aqui peguei os paramêtros da request

	paramID, err := strconv.ParseUint(params["id"], 10, 32) // Como o ID que vem via paramêtro via URL é string, precisamos converte-lo para int.

	if err != nil {
		response.Write([]byte("Erro ao converter paramêtro ID"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}

	if err = database.Get(db, response, paramID); err != nil {
		return
	}

	response.WriteHeader(http.StatusOK) // Informo que o Status Code vai ser (200)
}

// SearchAllHTTP :: Está função é chamada quando á uma request no "/api/v1/user" (\n)
// e ela retorna uma lista de Json com todos os usuários que existem no banco de dados. (Método GET).
func SearchAllHTTP(response http.ResponseWriter, request *http.Request) {
	db, err := database.Connect()

	if err != nil {
		response.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}

	if err := database.GetAll(db, response); err != nil {
		return
	}

	response.WriteHeader(http.StatusOK) // Informo que o Status Code vai ser (200)
}
