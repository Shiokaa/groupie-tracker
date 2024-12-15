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

	pageNum, _ := strconv.Atoi(page)

	if pageNum > 1 && pageNum < 10 {

		start := (pageNum - 1) * 20
		end := pageNum * 20

		data.Results = data.Results[start:end]

	}

	if pageNum < 1 && pageNum > 10 {
		data.Results = data.Results[0:20]
	}

	templates.Temp.ExecuteTemplate(w, "collection", data)

}
