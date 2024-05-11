package api

import (
	"net/http"
	"strconv"

	"github.com/blkzy/Curso-Go-Projetos/Database/2-CRUD/database"
	"github.com/gorilla/mux"
)

// DeleteHTTP :: This function is called when there is a request to "/api/v1/user/{ID}" (DELETE method).
func DeleteHTTP(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request) // Here I got the parameters from the request

	paramID, err := strconv.ParseUint(params["id"], 10, 32) // Since the ID that comes as a parameter in the URL is a string, we need to convert it to an int.

	if err != nil {
		response.Write([]byte("Error converting ID parameter!"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Write([]byte("Error connecting to the database!"))
		return
	}

	if err = database.Delete(db, response, paramID); err != nil {
		return
	}

	response.WriteHeader(http.StatusNoContent)
}
