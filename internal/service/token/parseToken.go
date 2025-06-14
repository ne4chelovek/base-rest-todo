package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
)

func (s *service) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	return claims.UserId, nil
}
