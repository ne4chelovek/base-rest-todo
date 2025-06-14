package item

import (
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"golang.org/x/net/context"
)

func (s *service) GetById(ctx context.Context, userId, itemId int) (*model.TodoItem, error) {
	return s.itemRepository.GetById(ctx, userId, itemId)
}
