package repository

import (
	"fmt"
	"github.com/KonstantinP85/api-go"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user api_go.User) (int, string, error) {
	var id int
	var email string
	query := fmt.Sprintf("INSERT INTO %s (email, password, is_active) values ($1, $2, $3) RETURNING id, email", usersTable)
	row := r.db.QueryRow(query, user.Email, user.Password, false)
	if err := row.Scan(&id, &email); err != nil {
		return 0, "", err
	}

	return id, email, nil
}

func (r *Auth) GetSignUser(email, password string) (api_go.User, error) {
	var user api_go.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email = $1 AND password = $2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}