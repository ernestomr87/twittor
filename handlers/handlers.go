package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ernestomr87/twittor/middlew"
	"github.com/ernestomr87/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers RUNNING GO SERVER
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDBConnection(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
