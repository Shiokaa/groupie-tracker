package collection

import (
	"Groupie-Tracker/api"
	"Groupie-Tracker/templates"
	"net/http"
)

func collectionController(w http.ResponseWriter, r *http.Request) {

	data := api.GetChar()

	templates.Temp.ExecuteTemplate(w, "collection", data)

}
