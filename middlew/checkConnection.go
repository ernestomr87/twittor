package middlew

import (
	"net/http"

	"github.com/ernestomr87/twittor/bd"
)

func CheckDBConnection(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckDBConnection() == 0 {
			http.Error(w, "Connection failed", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
