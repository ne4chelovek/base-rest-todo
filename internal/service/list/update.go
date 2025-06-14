package list

import (
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"golang.org/x/net/context"
)

func (s *service) Update(ctx context.Context, userId int, listId int, input *model.UpdateListInput) error {
	if err := input.Valid(); err != nil {
		return err
	}
	return s.listRepository.Update(ctx, userId, listId, input)
}
