package item

import (
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"golang.org/x/net/context"
)

func (s *service) GetAllItem(ctx context.Context, userId, listId int) ([]*model.TodoItem, error) {
	return s.itemRepository.GetAllItem(ctx, userId, listId)
}
