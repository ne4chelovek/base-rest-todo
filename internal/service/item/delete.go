package item

import "golang.org/x/net/context"

func (s *service) Delete(ctx context.Context, userId, itemId int) error {
	return s.itemRepository.Delete(ctx, userId, itemId)
}
