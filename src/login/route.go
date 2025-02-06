package login

import "net/http"

func LoginRouter() {
	http.HandleFunc("/login", loginController)
	http.HandleFunc("/login/traitement", loginTraitement)
}
