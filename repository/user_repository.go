package repository

import (
	"context"
	"github.com/dyeghocunha/golang-auth/db"
	"github.com/dyeghocunha/golang-auth/model"
)

func GetUserByEmail(email string) (*model.User, error) {
	query := `SELECT id, email, two_fa_secret, is_two_fa_enabled, created_at, updated_at,password_hash FROM users WHERE email = $1 LIMIT 1`
	row := db.Conn.QueryRow(context.Background(), query, email)

	var u model.User
	err := row.Scan(&u.ID, &u.Email, &u.TwoFASecret, &u.IsTwoFAEnabled, &u.CreatedAt, &u.UpdatedAt, &u.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func CreateUser(email, passwordHash string) error {
	query := `INSERT INTO users (email, password_hash, created_at, updated_at) VALUES ($1, $2, CURRENT_DATE, CURRENT_DATE)`
	_, err := db.Conn.Exec(context.Background(), query, email, passwordHash)
	return err
}

func UpdateUserTwoFA(email, secret string, enabled bool) error {
	query := `UPDATE users SET two_fa_secret=$1, is_two_fa_enabled=$2, updated_at=NOW() WHERE email=$3`
	_, err := db.Conn.Exec(context.Background(), query, secret, enabled, email)
	return err
}

func UpdateUserEmail(id int, newEmail string) error {
	query := `UPDATE users SET email =$1, updated_at=NOW() WHERE id=$2`
	_, err := db.Conn.Exec(context.Background(), query, newEmail, id)
	return err
}

func DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := db.Conn.Exec(context.Background(), query, id)
	return err
}

func Enable2Fa(email string) error {
	query := `UPDATE users SET is_two_fa_enabled = true WHERE email = $1`
	_, err := db.Conn.Exec(context.Background(), query, email)
	return err
}
