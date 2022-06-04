package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ernestomr87/twittor/bd"
	"github.com/ernestomr87/twittor/models"
)

// Register Insert User in DB
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if len(t.Email) < 6 {
		http.Error(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	_, found, _ := bd.CheckIfUserExists(t.Email)
	if found {
		http.Error(w, "User already exists", http.StatusNotAcceptable)
		return
	}

	_, status, err := bd.RegisterUser(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	if status == false {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
