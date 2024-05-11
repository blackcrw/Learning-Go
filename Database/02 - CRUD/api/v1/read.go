package api

import (
	"net/http"
	"strconv"

	"github.com/blkzy/Curso-Go-Projetos/Database/2-CRUD/database"
	"github.com/gorilla/mux"
)

// SearchHTTP :: This function will be called when there is a request with the path "/api/v1/user/{id}" (GET method).
func SearchHTTP(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request) // Here I'm getting the parameters from the request

	paramID, err := strconv.ParseUint(params["id"], 10, 32) // Since the ID that comes as a parameter in the URL is a string, we need to convert it to an int.

	if err != nil {
		response.Write([]byte("Error converting ID parameter"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Write([]byte("Error connecting to the database!"))
		return
	}

	if err = database.Get(db, response, paramID); err != nil {
		return
	}

	response.WriteHeader(http.StatusOK) // Informing that the Status Code will be (200)
}

// SearchAllHTTP :: This function is called when there is a request to "/api/v1/user" (GET method)
// and it returns a list of JSON with all the users in the database.
func SearchAllHTTP(response http.ResponseWriter, request *http.Request) {
	db, err := database.Connect()

	if err != nil {
		response.Write([]byte("Error connecting to the database!"))
		return
	}

	if err := database.GetAll(db, response); err != nil {
		return
	}

	response.WriteHeader(http.StatusOK) // Informing that the Status Code will be (200)
}
