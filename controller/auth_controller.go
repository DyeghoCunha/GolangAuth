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
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, `{"error":"User not found"}`, http.StatusUnauthorized)
		return
	}

	if !util.CheckPasswordHash(user.PasswordHash, req.Password) {
		http.Error(w, `{"error":"Invalid password"}`, http.StatusUnauthorized)
		return
	}

	if user.IsTwoFAEnabled {
		partialToken, err := util.GeneratePartialJWT(user.Email)
		if err != nil {
			http.Error(w, `{"error":"Error generating partial token"}`, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"partial_token": partialToken,
			"message":       "2FA required, please verify your code",
		})
		return
	}

	// 2FA não ativado → gerar access/refresh normal
	accessToken, err := util.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, `{"error":"Error generating access token"}`, http.StatusInternalServerError)
		return
	}

	refreshToken, err := util.GenerateRefreshToken(user.Email)
	if err != nil {
		http.Error(w, `{"error":"Error generating refresh token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"message":       "Login successful",
	})
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
	emailCtx := r.Context().Value("user_email")
	if emailCtx == nil {
		http.Error(w, "Email não encontrado no contexto", http.StatusUnauthorized)
		return
	}

	email := emailCtx.(string)

	secret, err := util.Generate2FASecret(email)
	if err != nil {
		http.Error(w, "Erro ao gerar o segredo 2FA: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = repository.UpdateUserTwoFA(email, secret, true)
	if err != nil {
		http.Error(w, "Erro ao atualizar 2FA: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"secret": secret,
	})
}

func GenerateQRCodeHandler(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("user_email")
	if email == nil {
		http.Error(w, "Email not found in token", http.StatusUnauthorized)
		return
	}
	emailStr := email.(string)

	user, err := repository.GetUserByEmail(emailStr)
	if err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}

	if user.TwoFASecret == "" {
		http.Error(w, "2FA is not enabled for this user", http.StatusBadRequest)
		return
	}

	otpKey, err := otp.NewKeyFromURL(fmt.Sprintf("otpauth://totp/GolangAuth:%s?secret=%s&issuer=GolangAuth", emailStr, user.TwoFASecret))
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
