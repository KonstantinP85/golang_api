package repository

import (
	"github.com/KonstantinP85/api-go"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user api_go.User) (int, string, error)
	GetSignUser(email, password string) (api_go.User, error)
}

type Book interface {
	GetBook(id int) (api_go.Book, error)
	GetBooks() ([]api_go.Book, error)
}

type User interface {
	GetUser(id int) (api_go.User, error)
	GetUsers() ([]api_go.User, error)
	UpdateUser(id int, input api_go.UpdateUserInput) error
	Patch(userId int, input api_go.PatchUserInput) error
	DeleteUser(id int) error
}

type Repository struct {
	Authorization
	Book
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
		User: NewUserDB(db),
		Book: NewBookDB(db),
	}
}
