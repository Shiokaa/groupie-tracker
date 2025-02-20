package login

import (
	"GroupieTracker/modele"
	"GroupieTracker/templates"
	"net/http"
	"strings"
)

var IsRegistered bool
var UserConnected string
var Id int

type Data struct {
	IsRegistered  bool
	UserConnected string
	Err           string
}

func loginController(w http.ResponseWriter, r *http.Request) {
	var data Data

	data.IsRegistered = IsRegistered
	data.UserConnected = UserConnected

	if r.FormValue("disconnect") == "true" {
		IsRegistered = false
		UserConnected = ""
		Id = 0
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	errMsg := r.FormValue("err")
	if errMsg != "" {
		data.Err = getErrorMessage(errMsg)
	}

	templates.Temp.ExecuteTemplate(w, "login", data)
}

func getErrorMessage(code string) string {
	switch code {
	case "champvide":
		return "Veuillez remplir tous les champs"
	case "mdp":
		return "Mot de passe incorrect"
	case "notfound":
		return "Utilisateur introuvable"
	default:
		return "Une erreur est survenue"
	}
}

func loginTraitement(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))

	if username == "" || password == "" {
		http.Redirect(w, r, "/login?err=champvide", http.StatusSeeOther)
		return
	}

	users := modele.JsonRead()

	for _, elem := range users {
		if strings.EqualFold(elem.Username, username) {
			if elem.Password == password {
				IsRegistered = true
				UserConnected = username
				Id = elem.Id
				http.Redirect(w, r, "/home", http.StatusSeeOther)
				return
			} else {
				http.Redirect(w, r, "/login?err=mdp", http.StatusSeeOther)
				return
			}
		}
	}

	http.Redirect(w, r, "/login?err=notfound", http.StatusSeeOther)
}
