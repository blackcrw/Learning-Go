package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/blakcrw/learning-go/Database/02 - CRUD/database"
)

// CreateHTTP :: This function is called when there is a request to "/api/v1/user" (POST method).
func CreateHTTP(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println(err)
		response.Write([]byte("Failed to read request body!"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Write([]byte("Error connecting to the database!"))
		return
	}

	idUser, err := database.Create(db, response, body)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte(fmt.Sprintf("User created successfully! ID: %d", idUser)))
}
