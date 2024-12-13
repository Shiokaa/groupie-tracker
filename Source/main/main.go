package main

import (
	"Groupie-Tracker/collection"
	"Groupie-Tracker/home"
	"Groupie-Tracker/research"
	"Groupie-Tracker/templates"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	templates.InitTemplates()

	home.HomeRouter()
	collection.CollectionRouter()
	research.ResearchRouter()

	http.ListenAndServe(":8080", nil)

}
