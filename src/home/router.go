package home

import "net/http"

func HomeRouter() {
	http.HandleFunc("/home", homeController)
}
