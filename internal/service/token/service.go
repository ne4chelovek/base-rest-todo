package token

import (
	"github.com/ne4chelovek/base-rest-todo/internal/repository"
	"time"
)

const (
	salt       = "afpodkh[apdfkg"
	signingKey = "aopk214@!$9dfgpaksdg"
	tokenTTL   = 12 * time.Hour
)

type service struct {
	authRepository repository.Authorization
}

func NewService(authRepository repository.Authorization) *service {
	return &service{
		authRepository: authRepository,
	}
}
