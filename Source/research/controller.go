package research

import (
	"Groupie-Tracker/api"
	"Groupie-Tracker/templates"
	"net/http"
	"strconv"
	"strings"
)

type apiData struct {
	Episode   api.Episode
	Location  api.Location
	Character api.Character
	GetById   bool
	GetByName bool
	Id        string
	Name      string
}

func researchController(w http.ResponseWriter, r *http.Request) {

	episode := api.GetEpisode()
	location := api.GetLocation()
	character := api.GetCharacters()
	input := r.FormValue("research")

	var data apiData

	_, err := strconv.Atoi(input)

	if err == nil {

		for _, elems := range episode.Results {
			if strconv.Itoa(elems.Id) == input {
				data.Episode.Results = append(data.Episode.Results, elems)
				break
			}
		}

		for _, elems := range location.Results {
			if strconv.Itoa(elems.Id) == input {
				data.Location.Results = append(data.Location.Results, elems)
				break
			}
		}

		for _, elems := range character.Results {
			if strconv.Itoa(elems.Id) == input {
				data.Character.Results = append(data.Character.Results, elems)
				break
			}
		}

		data.GetById = true
		data.GetByName = false
		data.Id = input
	}

	if err != nil {

		for _, elems := range episode.Results {
			if strings.Contains(strings.ToLower(elems.Name), strings.ToLower(input)) {
				data.Episode.Results = append(data.Episode.Results, elems)
				break
			}
		}

		for _, elems := range location.Results {
			if strings.Contains(strings.ToLower(elems.Name), strings.ToLower(input)) {
				data.Location.Results = append(data.Location.Results, elems)
				break
			}
		}

		for _, elems := range character.Results {
			if strings.Contains(strings.ToLower(elems.Name), strings.ToLower(input)) {
				data.Character.Results = append(data.Character.Results, elems)
				break
			}
		}
		data.GetById = false
		data.GetByName = true
		data.Name = input
	}

	templates.Temp.ExecuteTemplate(w, "research", data)
}
