package middleWear

import "net/http"

func (s *ApiDbMiddleWear) EnableCors(next http.HandlerFunc)http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // lub "*" dla wszystkich originów
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        next.ServeHTTP(w, r)
    }
}