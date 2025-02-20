package favorite

import (
	"GroupieTracker/api"
	"GroupieTracker/login"
	"GroupieTracker/modele"
	"GroupieTracker/templates"
	"net/http"
	"strconv"
)

type UserDetails struct {
	IsRegistered  bool
	UserConnected string
	NoFavorite    bool
	Favorite      []modele.Favorite
}

func favoriteController(w http.ResponseWriter, r *http.Request) {
	var data UserDetails

	data.UserConnected = login.UserConnected
	data.IsRegistered = login.IsRegistered

	if data.IsRegistered {
		file := modele.JsonRead()
		for _, elem := range file {
			if elem.Id == login.Id {
				data.Favorite = elem.Favorite
				if len(data.Favorite) == 0 {
					data.NoFavorite = true
				}
			}
		}
	}

	templates.Temp.ExecuteTemplate(w, "favorite", data)
}

func favoriteTraitement(w http.ResponseWriter, r *http.Request) {
	delete := r.FormValue("delete")
	charId := r.FormValue("id")

	charIdNumber, err := strconv.Atoi(charId)
	if err != nil {
		http.Redirect(w, r, "/favorite", http.StatusSeeOther)
		return
	}

	if delete != "" {
		modele.FavoriteDelete(login.Id, charIdNumber)
		http.Redirect(w, r, "/favorite", http.StatusSeeOther)
		return
	}

	character := api.GetCharacterById(charId)

	if charId == "" {
		http.Redirect(w, r, "/favorite", http.StatusSeeOther)
		return
	}

	favorite := modele.Favorite{Id: character.Id, Name: character.Name, Status: character.Status, Species: character.Species, Gender: character.Gender, Images: character.Images}

	modele.FavoriteAdd(favorite, login.Id)

	http.Redirect(w, r, "/favorite", http.StatusSeeOther)
}
