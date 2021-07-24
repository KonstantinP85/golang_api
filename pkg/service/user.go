package service

import (
	"github.com/KonstantinP85/api-go"
	"github.com/KonstantinP85/api-go/pkg/repository"
)

type UserService struct {
	 repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (api_go.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserService) GetUsers() ([]api_go.User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) UpdateUser(id int, input api_go.UpdateUserInput) error {
	input.Password = generatePasswordHash(input.Password)
	return s.repo.UpdateUser(id, input)
}

func (s *UserService) Patch(id int, input api_go.PatchUserInput) error {
	return s.repo.Patch(id, input)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}