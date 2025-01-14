package research

import "net/http"

func ResearchRouter() {
	http.HandleFunc("/research", researchController)
}
