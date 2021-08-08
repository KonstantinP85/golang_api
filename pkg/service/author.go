package service

import (
	api_go "github.com/KonstantinP85/api-go"
	"github.com/KonstantinP85/api-go/pkg/repository"
)

type AuthorService struct {
	repo repository.Author
}

func NewAuthorService(repo repository.Author) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) CreateAuthor(author api_go.AuthorInput) (int, error) {
	return s.repo.CreateAuthor(author)
}

func (s *AuthorService) GetAuthors() ([]api_go.Author, error) {
	return s.repo.GetAuthors()
}

func (s *AuthorService) GetAuthor(id int) (api_go.AuthorResponse, error) {
	return s.repo.GetAuthor(id)
}