package service

import "github.com/KonstantinP85/api-go/pkg/repository"

type Authorization interface {

}

type Book interface {

}

type Service struct {
	Authorization
	Book
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}