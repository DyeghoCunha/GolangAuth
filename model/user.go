package model

type User struct {
	ID             int
	Email          string
	TwoFASecret    string
	IsTwoFAEnabled bool
	CreatedAt      string
	UpdatedAt      string
}
