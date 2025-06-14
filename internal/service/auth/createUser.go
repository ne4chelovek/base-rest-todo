package auth

import (
	"context"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(ctx context.Context, user *model.User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("Слабый пароль  %v", err)
		return 0, err
	}

	user.Password = string(hashedPassword)

	id, err := s.authRepository.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
