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
	GetBook(id int) (api_go.Book, error)
	GetBooks() ([]api_go.Book, error)
}

type Service struct {
	Authorization
	User
	Book
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		User: NewUserService(repo.User),
		Book: NewBookService(repo.Book),
	}
}