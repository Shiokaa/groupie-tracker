package collection

import "net/http"

func CollectionRouter() {
	http.HandleFunc("/collection", collectionController)
}
