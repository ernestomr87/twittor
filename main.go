package main

import (
	"log"

	"github.com/ernestomr87/twittor/bd"
	"github.com/ernestomr87/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("WithoutConnection DB")
		return
	}
	handlers.Handlers()
}
