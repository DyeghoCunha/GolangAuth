package controller

import (
	"encoding/json"
	"fmt"
	"github.com/dyeghocunha/golang-auth/repository"
	"github.com/dyeghocunha/golang-auth/util"
	"github.com/pquerna/otp"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	user, err := repository.GetUserByEmail(req.Email)

	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	if !util.CheckPasswordHash(user.PasswordHash, req.Password) {
		http.Error(w, "Hash and Password doesen´t match", http.StatusUnauthorized)
		return
	}

	if user.IsTwoFAEnabled {
		partialToken, err := util.GeneratePartialJWT(user.Email)
		if err != nil {
			http.Error(w, "Error generating partial token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"partial_token": partialToken,
			"message":       "2FA required, please verify your code",
		})
		return
		accessToken, err := util.GenerateJWT(user.Email)
		if err != nil {
			http.Error(w, "Erro ao gerar access token", http.StatusInternalServerError)
			return
		}

		refreshToken, err := util.GenerateRefreshToken(user.Email)
		if err != nil {
			http.Error(w, "Erro ao gerar refresh token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}

}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid E-mail", http.StatusBadRequest)
		return
	}
	password := req.Password
	//hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hash, err := util.HashPassword(password)
	if err != nil {
		http.Error(w, "Erro ao criar hash da senha", http.StatusInternalServerError)
		return
	}

	err = repository.CreateUser(req.Email, hash)
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

func Enable2FAHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	secret, err := util.Generate2FASecret(email)
	if err != nil {
		http.Error(w, "Error generating 2FA secret: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = repository.UpdateUserTwoFA(email, secret, true)
	if err != nil {
		http.Error(w, "Error updating 2FA: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"secret": secret,
	})
}

func GenerateQRCodeHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}
	if user.TwoFASecret == "" {
		http.Error(w, "2FA is not enabled for this user", http.StatusBadRequest)
		return
	}
	otpKey, err := otp.NewKeyFromURL(fmt.Sprintf("otpauth://totp/GolangAuth:%s?secret=%s&issuer=GolangAuth", email, user.TwoFASecret))
	if err != nil {
		http.Error(w, "Error generating OTP key: "+err.Error(), http.StatusInternalServerError)
		return
	}
	png, err := qrcode.Encode(otpKey.URL(), qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Error generating QR code: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}
