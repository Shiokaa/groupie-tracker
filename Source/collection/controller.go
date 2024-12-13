package collection

import (
	"Groupie-Tracker/api"
	"Groupie-Tracker/templates"
	"net/http"
	"strconv"
)

func collectionController(w http.ResponseWriter, r *http.Request) {

	data := api.GetCharacters()
	data.Results = data.Results[0:20]
 
	page := r.FormValue("page")

	if page != "" {

		pageNum, _ := strconv.Atoi(page)

		start := (pageNum - 1) * 20
		end := pageNum * 20

		data.Results = data.Results[start:end]

	}

	templates.Temp.ExecuteTemplate(w, "collection", data)

}
