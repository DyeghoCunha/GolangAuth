package controller

import (
	"encoding/json"
	"github.com/dyeghocunha/golang-auth/repository"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Error retrieving user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
