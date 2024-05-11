package api

import (
	"net/http"
	"strconv"

	"github.com/blkzy/Curso-Go-Projetos/Database/2-CRUD/database"
	"github.com/gorilla/mux"
)

// DeleteHTTP :: Está função é chamada quando ah um request no "/api/v1/user/{ID}" (Método DELETE).
func DeleteHTTP(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request) // Aqui peguei os paramêtros da request

	paramID, err := strconv.ParseUint(params["id"], 10, 32) // Como o ID que vem via paramêtro via URL é string, precisamos converte-lo para int.

	if err != nil {
		response.Write([]byte("Erro ao converter paramêtro ID!"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}

	if err = database.Delete(db, response, paramID); err != nil {
		return
	}

	response.WriteHeader(http.StatusNoContent)
}
