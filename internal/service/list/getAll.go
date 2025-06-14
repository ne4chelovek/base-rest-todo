package list

import (
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"golang.org/x/net/context"
)

func (s *service) GetAll(ctx context.Context, userId int) ([]*model.TodoList, error) {
	return s.listRepository.GetAll(ctx, userId)
}
