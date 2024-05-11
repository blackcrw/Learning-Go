package api

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/blkzy/Curso-Go-Projetos/Database/2-CRUD/database"
	"github.com/gorilla/mux"
)

// UpdateHTTP :: Está função é chamada quando ah uma request no "/api/v1/user/{ID}" (Método PUT).
func UpdateHTTP(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request) // Aqui peguei os paramêtros da request

	paramID, err := strconv.ParseUint(params["id"], 10, 32) // Como o ID que vem via paramêtro via URL é string, precisamos converte-lo para int.

	if err != nil {
		response.Write([]byte("Erro ao converter paramêtro ID!"))
		return
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		response.Write([]byte("Erro ao ler o corpo da requisição."))
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}

	if err = database.Update(db, response, body, paramID); err != nil {
		return
	}

	response.WriteHeader(http.StatusOK) // Informo que o Status Code vai ser (200)

}
