package collection

import (
	"GroupieTracker/api"
	"GroupieTracker/login"
	"GroupieTracker/templates"
	"fmt"
	"net/http"
	"strconv"
)

type DataStruct struct {
	Characters   api.Character
	IsRegistered bool
	Filter       struct {
		Status  string
		Gender  string
		Counter int
	}
	Pagination struct {
		Page int
		Prev string
		Next string
	}
}

func formatPagination(page int, data *DataStruct) {
	link := "/collection?"

	if data.Filter.Status != "" {
		link = fmt.Sprintf("%sstatus=%s", link, data.Filter.Status)
	}

	if data.Filter.Gender != "" {
		link = fmt.Sprintf("%s&gender=%s", link, data.Filter.Gender)
	}

	data.Pagination.Next = fmt.Sprintf("%s&page=%d", link, page+1)
	data.Pagination.Prev = fmt.Sprintf("%s&page=%d", link, page-1)
}

func collectionCut(pageNum int, data *DataStruct) {
	// Nombre d'éléments par page
	const itemsPerPage = 20
	totalItems := len(data.Characters.Results)

	// Déterminer les indices de début et de fin
	start := (pageNum - 1) * itemsPerPage
	end := start + itemsPerPage

	// Validation des limites
	if start >= totalItems {
		// Si la page est hors limites, retourner une liste vide
		data.Characters.Results = nil
		return
	}

	if end > totalItems {
		end = totalItems // Limiter l'indice de fin au nombre total d'éléments
	}

	// Découper les résultats selon les indices calculés
	data.Characters.Results = data.Characters.Results[start:end]
}

func collectionController(w http.ResponseWriter, r *http.Request) {

	var data DataStruct
	data.IsRegistered = login.IsRegistered

	reset := r.FormValue("reset")
	page := r.FormValue("page")
	pageInt, pageErr := strconv.Atoi(page)

	if pageErr != nil {
		pageInt = 1
		data.Pagination.Page = pageInt
	} else {
		data.Pagination.Page = pageInt
	}

	data.Filter.Status = r.FormValue("status")
	data.Filter.Gender = r.FormValue("gender")

	if reset != "" {
		data.Filter.Status = ""
		data.Filter.Gender = ""
	}

	if data.Filter.Status != "" && data.Filter.Gender == "" {
		for _, elem := range api.GetAllCharacters().Results {
			if elem.Status == data.Filter.Status {
				data.Characters.Results = append(data.Characters.Results, elem)
			}
		}
		data.Filter.Counter = 1
	}

	if data.Filter.Status == "" && data.Filter.Gender != "" {
		for _, elem := range api.GetAllCharacters().Results {
			if elem.Gender == data.Filter.Gender {
				data.Characters.Results = append(data.Characters.Results, elem)
			}
		}
		data.Filter.Counter = 1
	}

	if data.Filter.Status != "" && data.Filter.Gender != "" {
		for _, elem := range api.GetAllCharacters().Results {
			if elem.Status == data.Filter.Status && elem.Gender == data.Filter.Gender {
				data.Characters.Results = append(data.Characters.Results, elem)
			}
		}
		data.Filter.Counter = 2
	}

	if data.Filter.Status == "" && data.Filter.Gender == "" {
		data.Characters = api.GetCharacters()
	}

	collectionCut(data.Pagination.Page, &data)
	formatPagination(data.Pagination.Page, &data)

	templates.Temp.ExecuteTemplate(w, "collection", data)

}
