package login

import (
	"GroupieTracker/templates"
	"net/http"
)

var IsRegistered bool
var UserConnected string
var Id int

func loginController(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "login", nil)
}

func loginTraitement(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("username") == "" || r.FormValue("password") == "" {
		http.Redirect(w, r, "/login?error=champvide", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
