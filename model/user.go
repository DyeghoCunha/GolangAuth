package model

type User struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	PasswordHash   string `json:"password_hash"`
	TwoFASecret    string `json:"two_fa_secret"`
	IsTwoFAEnabled bool   `json:"is_two_fa_enabled"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
