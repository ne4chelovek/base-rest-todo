package list

import "golang.org/x/net/context"

func (s *service) Delete(ctx context.Context, userId int, listId int) error {
	return s.listRepository.Delete(ctx, userId, listId)
}
