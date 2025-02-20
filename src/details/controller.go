package details

import (
	"GroupieTracker/api"
	"GroupieTracker/login"
	"GroupieTracker/templates"
	"net/http"
)

type CharacterDetails struct {
	Character     api.SingleCharacter
	IsRegistered  bool
	UserConnected string
	IsEmpty       bool
}

func detailsController(w http.ResponseWriter, r *http.Request) {
	var data CharacterDetails

	data.UserConnected = login.UserConnected
	data.IsRegistered = login.IsRegistered
	data.Character = api.GetCharacterById(r.FormValue("id"))

	if data.Character == (api.SingleCharacter{}) {
		data.IsEmpty = true
	}

	templates.Temp.ExecuteTemplate(w, "details", data)
}
