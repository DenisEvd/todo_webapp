package service

import (
	"todo_webapp"
	"todo_webapp/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_webapp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Task interface {
}

type Service struct {
	Authorization
	Task
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep),
	}
}
