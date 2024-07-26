package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/rtzgod/todo-app"
	"github.com/rtzgod/todo-app/internal/repository"
)

const salt = "dfofskdpofk33k24k9234k023jf0238fj0293jf293"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
