package main

import (
	"Groupie-Tracker/home"
	"Groupie-Tracker/templates"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	templates.InitTemplates()

	home.HomeRouter()

	http.ListenAndServe(":8080", nil)
}
