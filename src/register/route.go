package register

import "net/http"

func RegisterRouter() {
	http.HandleFunc("/register/traitement", registerTraitement)
	http.HandleFunc("/register", registerController)
}
