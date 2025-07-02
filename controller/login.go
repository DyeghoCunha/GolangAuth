package controller

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Email == "" || req.Password == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	if !util.CheckPasswordHash(req.Password, user.PasswordHash) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	if user.IsTwoFAEnabled {
		partialToken, _ := util.GeneratePartialJWT(req.Email)
		w.WriteHeader(http.StatusPartialContent)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "2FA required",
			"token":   partialToken,
		})
		return
	}
	token, _ := util.GenerateJWT(req.Email)
	refresh, _ := util.GenerateRefreshToken(req.Email)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  token,
		"refresh_token": refresh,
	})
}
*/
