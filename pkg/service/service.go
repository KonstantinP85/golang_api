package service

import (
	"github.com/KonstantinP85/api-go"
	"github.com/KonstantinP85/api-go/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mocks.go

type Authorization interface {
	CreateUser(user api_go.User) (int, string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetUser(id int) (api_go.User, error)
	GetUsers() ([]api_go.User, error)
	UpdateUser(id int, input api_go.UpdateUserInput) error
	Patch(id int, input api_go.PatchUserInput) error
	DeleteUser(id int) error
}

type Book interface {
	GetBook(id int) (api_go.BookResponse, error)
	GetBooks(authorId int) ([]api_go.BookResponse, error)
	CreateBook(book api_go.BookInput) (int, error)
}

type Author interface {
	CreateAuthor(author api_go.AuthorInput) (int, error)
	GetAuthors() ([]api_go.Author, error)
	GetAuthor(id int) (api_go.AuthorResponse, error)
}

type Service struct {
	Authorization
	User
	Book
	Author
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		User: NewUserService(repo.User),
		Book: NewBookService(repo.Book),
		Author: NewAuthorService(repo.Author),
	}
}