package service

import (
	"github.com/KonstantinP85/api-go"
	"github.com/KonstantinP85/api-go/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBook(id int) (api_go.BookResponse, error) {
	return s.repo.GetBook(id)
}

func (s *BookService) GetBooks(authorId int) ([]api_go.BookResponse, error) {
	return s.repo.GetBooks(authorId)
}

func (s *BookService) CreateBook(book api_go.BookInput) (int, error) {
	return s.repo.CreateBook(book)
}

