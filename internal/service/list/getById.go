package list

import (
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"golang.org/x/net/context"
)

func (s *service) GetById(ctx context.Context, userId int, listId int) (*model.TodoList, error) {
	return s.listRepository.GetById(ctx, userId, listId)
}
