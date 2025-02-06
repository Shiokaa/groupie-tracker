package register

import (
	"GroupieTracker/modele"
	"GroupieTracker/templates"
	"net/http"
)

func registerController(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("exist") == "username" {
		templates.Temp.ExecuteTemplate(w, "register", "Username already exist")
		return
	}

	templates.Temp.ExecuteTemplate(w, "register", nil)
}

func registerTraitement(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	users := modele.JsonRead()

	for _, elem := range users {
		if elem.Username == username {
			http.Redirect(w, r, "/register?exist=username", http.StatusSeeOther)
			return
		}
	}

	newUser := modele.User{Username: username, Password: password}

	modele.JsonWrite(newUser)

	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
