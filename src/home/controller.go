package home

import (
	"GroupieTracker/login"
	"GroupieTracker/templates"
	"net/http"
)

func homeController(w http.ResponseWriter, r *http.Request) {
	isRegistered := login.IsRegistered

	templates.Temp.ExecuteTemplate(w, "home", isRegistered)
}
