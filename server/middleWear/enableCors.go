package middleWear

import "net/http"

func EnableCors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // lub "*" dla wszystkich originów
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        next.ServeHTTP(w, r)
    })
}