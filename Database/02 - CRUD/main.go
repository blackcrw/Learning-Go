package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/blkzy/Curso-Go-Projetos/Database/2-CRUD/api/v1" // The routes will call the functions from this package.
	"github.com/gorilla/mux"
)

func init() {
	fmt.Println(`
			  .__          .__          
  ______ ____ |__|         |  | _____   
 /  ___// __ \|  |  ______ |  | \__  \  
 \___ \\  ___/|  | /_____/ |  |__/ __ \_
/____  >\___  >__|         |____(____  /
	 \/     \/                       \/ 	`)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user", api.CreateHTTP).Methods(http.MethodPost)        // This is the route that creates a user in the database.
	router.HandleFunc("/api/v1/user", api.SearchAllHTTP).Methods(http.MethodGet)      // This is the route that retrieves all users from the database.
	router.HandleFunc("/api/v1/user/{id}", api.SearchHTTP).Methods(http.MethodGet)    // This is the route that retrieves users by id from the database.
	router.HandleFunc("/api/v1/user/{id}", api.UpdateHTTP).Methods(http.MethodPut)    // This is the route that updates user data such as Name and Email.
	router.HandleFunc("/api/v1/user/{id}", api.DeleteHTTP).Methods(http.MethodDelete) // This is the route that deletes the user from the database.

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
