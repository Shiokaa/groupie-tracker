package collection

import (
	"Groupie-Tracker/api"
	"Groupie-Tracker/templates"
	"net/http"
)

func collectionController(w http.ResponseWriter, r *http.Request) {

	data := api.GetCharacters()

	data.Results = data.Results[0:20]

	templates.Temp.ExecuteTemplate(w, "collection", data)

}
