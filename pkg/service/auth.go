package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"todo_webapp"
	"todo_webapp/pkg/repository"
)

const (
	salt       = "h234hh234jh4hkj545jg"
	tokenTTL   = 12 * time.Hour
	signingKey = "sfkgsksdgkdsssdpafcpvxck"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

func (a *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := a.repo.GetUser(username, a.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
