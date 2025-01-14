package home

import (
	"Groupie-Tracker/templates"
	"net/http"
)

func homeController(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "home", nil)
}
