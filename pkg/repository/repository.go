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
	GetBook(id int) (api_go.BookResponse, error)
	GetBooks(authorId int) ([]api_go.BookResponse, error)
	CreateBook(book api_go.BookInput) (int, error)
}

type User interface {
	GetUser(id int) (api_go.User, error)
	GetUsers() ([]api_go.User, error)
	UpdateUser(id int, input api_go.UpdateUserInput) error
	Patch(userId int, input api_go.PatchUserInput) error
	DeleteUser(id int) error
}

type Author interface {
	CreateAuthor(author api_go.AuthorInput) (int, error)
	GetAuthors() ([]api_go.Author, error)
	GetAuthor(id int) (api_go.AuthorResponse, error)
}

type Repository struct {
	Authorization
	Book
	User
	Author
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
		User: NewUserDB(db),
		Book: NewBookDB(db),
		Author: NewAuthorDB(db),
	}
}
