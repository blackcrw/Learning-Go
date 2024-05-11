package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type users struct {
	Name  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("Template/*.html"))

	http.HandleFunc("/index", func(response http.ResponseWriter, request *http.Request) {
		infos := users{Name: "Blkz", Email: "teste@gmail.com"}
		templates.ExecuteTemplate(response, "index.html", infos)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
