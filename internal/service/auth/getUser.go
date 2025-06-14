package auth

import (
	"context"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
)

func (s *service) GetUser(ctx context.Context, username string) (*model.UserInfo, error) {
	user, err := s.authRepository.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
