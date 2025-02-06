package favorite

import (
	"GroupieTracker/login"
	"GroupieTracker/templates"
	"net/http"
)

type Data struct {
	IsRegistered  bool
	UserConnected string
}

func favoriteController(w http.ResponseWriter, r *http.Request) {
	var data Data

	data.IsRegistered = login.IsRegistered
	data.UserConnected = login.UserConnected

	templates.Temp.ExecuteTemplate(w, "favorite", data)
}

func favoriteTraitement(w http.ResponseWriter, r *http.Request) {


	
	http.Redirect(w, r, "/favorite", http.StatusSeeOther)
}
