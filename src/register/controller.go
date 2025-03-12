package register

import (
	"GroupieTracker/login"
	"GroupieTracker/modele"
	"GroupieTracker/templates"
	"net/http"
	"regexp"
	"strings"
)

type Data struct {
	IsRegistered  bool
	UserConnected string
	Err           string
}

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,16}$`)
var passwordMinLength = 6

func registerController(w http.ResponseWriter, r *http.Request) {
	var data Data
	data.IsRegistered = login.IsRegistered
	data.UserConnected = login.UserConnected

	errMsg := r.FormValue("err")
	if errMsg != "" {
		data.Err = errMsg
	}

	if data.IsRegistered {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	templates.Temp.ExecuteTemplate(w, "register", data)
}

func registerTraitement(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))

	// Vérification des champs vides
	if username == "" || password == "" {
		http.Redirect(w, r, "/register?err=Veuillez remplir tous les champs", http.StatusSeeOther)
		return
	}

	// Vérification du format du nom d'utilisateur
	if !usernameRegex.MatchString(username) {
		http.Redirect(w, r, "/register?err=Nom d'utilisateur invalide (3-16 caractères, lettres, chiffres, tirets)", http.StatusSeeOther)
		return
	}

	// Vérification de la longueur du mot de passe
	if len(password) < passwordMinLength {
		http.Redirect(w, r, "/register?err=Mot de passe trop court (minimum 6 caractères)", http.StatusSeeOther)
		return
	}

	users := modele.JsonRead()

	// Vérification si l'utilisateur existe déjà
	for _, elem := range users {
		if strings.EqualFold(elem.Username, username) {
			http.Redirect(w, r, "/register?err=Nom d'utilisateur déjà pris", http.StatusSeeOther)
			return
		}
	}

	// Création et enregistrement du nouvel utilisateur
	newUser := modele.User{Username: username, Password: password}
	modele.JsonWrite(newUser)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
