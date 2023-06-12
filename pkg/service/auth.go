package service

import (
	"crypto/sha1"
	"fmt"
	"todo_webapp"
	"todo_webapp/pkg/repository"
)

const salt = "h234hh234jh4hkj545jg"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{repo: rep}
}

func (a *AuthService) CreateUser(user todo_webapp.User) (int, error) {
	user.Password = a.generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func (a *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%s", hash.Sum([]byte(salt)))
}
