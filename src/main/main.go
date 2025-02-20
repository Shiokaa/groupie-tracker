package main

import (
	"GroupieTracker/collection"
	"GroupieTracker/details"
	"GroupieTracker/favorite"
	"GroupieTracker/home"
	"GroupieTracker/login"
	"GroupieTracker/register"
	"GroupieTracker/research"
	"GroupieTracker/templates"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	templates.InitTemplates()

	home.HomeRouter()
	collection.CollectionRouter()
	research.ResearchRouter()
	register.RegisterRouter()
	favorite.FavoriteRouter()
	login.LoginRouter()
	details.DetailsRouter()

	http.ListenAndServe(":3000", nil)

}
