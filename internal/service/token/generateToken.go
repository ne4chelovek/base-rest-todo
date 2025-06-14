package token

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s service) GenerateToken(ctx context.Context, username string, password string) (string, error) {
	user, err := s.authRepository.GetUser(ctx, username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}
