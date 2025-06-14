package token

import (
	"github.com/ne4chelovek/base-rest-todo/internal/repository"
)

const (
	signingKey = "aopk214@!$9dfgpaksdg"
)

type service struct {
	authRepository repository.Authorization
}

func NewService(authRepository repository.Authorization) *service {
	return &service{
		authRepository: authRepository,
	}
}
