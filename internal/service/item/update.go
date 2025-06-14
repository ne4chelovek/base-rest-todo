package item

import (
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"golang.org/x/net/context"
)

func (s *service) Update(ctx context.Context, userId, itemId int, input *model.UpdateItemInput) error {
	return s.itemRepository.Update(ctx, userId, itemId, input)
}
