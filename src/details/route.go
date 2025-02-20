package details

import "net/http"

func DetailsRouter() {
	http.HandleFunc("/details", detailsController)
}
