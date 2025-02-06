package favorite

import "net/http"

func FavoriteRouter() {

	http.HandleFunc("/favorite", favoriteController)
	http.HandleFunc("/favorite/traitement", favoriteTraitement)
}
