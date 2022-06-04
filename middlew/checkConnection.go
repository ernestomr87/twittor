package middlew

import (
	"net/http"

	"github.com/ernestomr87/twittor/bd"
)

//	CheckDBConnection Middleware for testing status of db connection
func CheckDBConnection(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckDBConnection() == 0 {
			http.Error(w, "Connection DB failed", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
