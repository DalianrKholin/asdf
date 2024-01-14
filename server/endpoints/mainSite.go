package server

import "net/http"

func (s *ApiDbEndpoints) MainSite(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/entry.html")
}
