package item

import (
	"github.com/jackc/pgx/v5"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"golang.org/x/net/context"
)

func (s *service) Create(ctx context.Context, userId, listId int, item *model.TodoItem) (int, error) {
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	txRepoItem := s.itemRepository.WithTx(tx)

	_, err = s.listRepository.GetById(ctx, userId, listId)
	if err != nil {
		return 0, err
	}

	id, err := txRepoItem.Create(ctx, listId, item)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, err
	}

	return id, nil
}
