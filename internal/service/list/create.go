package list

import (
	"github.com/jackc/pgx/v5"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"golang.org/x/net/context"
)

func (s *service) Create(ctx context.Context, userId int, list *model.TodoList) (int, error) {
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	txRepo := s.listRepository.WithTx(tx)

	id, err := txRepo.Create(ctx, userId, list)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return id, nil
}
