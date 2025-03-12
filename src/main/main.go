package main

import (
	"GroupieTracker/about"
	"GroupieTracker/collection"
	"GroupieTracker/details"
	"GroupieTracker/favorite"
	"GroupieTracker/home"
	"GroupieTracker/login"
	"GroupieTracker/register"
	"GroupieTracker/research"
	"GroupieTracker/templates"
	"fmt"
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
	about.AboutRouter()

	fmt.Println("Serveur lancé sur le port 3000 : http://localhost:3000/home ")

	http.ListenAndServe(":3000", nil)

}
