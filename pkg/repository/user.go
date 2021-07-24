package repository

import (
	"fmt"
	"github.com/KonstantinP85/api-go"
	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserDB {
	return &UserDB{db: db}
}

func (r *UserDB) GetUser(id int) (api_go.User, error) {
	var user api_go.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UserDB) GetUsers() ([]api_go.User, error) {
	var users []api_go.User

	query := fmt.Sprintf("SELECT * FROM %s", usersTable)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UserDB) UpdateUser(id int, input api_go.UpdateUserInput) error {
	fmt.Println(input.Email)
	query:= fmt.Sprintf("UPDATE %s SET email=$1, password=$2  WHERE id=$3", usersTable)

	_, err := r.db.Exec(query, input.Email, input.Password, id)

	return err
}

func (r *UserDB) Patch(id int, input api_go.PatchUserInput) error {

	query := fmt.Sprintf("UPDATE %s SET is_active=$1 WHERE id=$2", usersTable)

	_, err := r.db.Exec(query, input.IsActive, id)

	return err
}

func (r *UserDB) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)

	_, err := r.db.Exec(query, id)

	return err
}