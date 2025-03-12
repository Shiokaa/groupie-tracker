package about

import "net/http"

func AboutRouter() {
	http.HandleFunc("/about", aboutController)
}
