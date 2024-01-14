package middleWear

import "net/http"

func (s *ApiDbMiddleWear) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO : to be Done
		next(w, r)
	}
}
