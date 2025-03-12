package about

import (
	"GroupieTracker/login"
	"GroupieTracker/templates"
	"net/http"
)

type Data struct {
	IsRegistered bool
	File         string
}

func aboutController(w http.ResponseWriter, r *http.Request) {

	var data Data

	data.IsRegistered = login.IsRegistered

	templates.Temp.ExecuteTemplate(w, "about", data)
}
