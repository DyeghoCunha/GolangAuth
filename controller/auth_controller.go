package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dyeghocunha/golang-auth/repository"
)

type RegisterRequest struct {
	Email string `json:"email"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid E-mail", http.StatusBadRequest)
		return
	}
	err = repository.CreateUser(req.Email)
	if err != nil {
		log.Println("Error creating user:", err)
		http.Error(w, "Error creating user"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}
