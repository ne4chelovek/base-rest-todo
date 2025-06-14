package list

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ne4chelovek/base-rest-todo/internal/repository"
)

type service struct {
	listRepository repository.TodoList
	dbPool         *pgxpool.Pool
}

func NewService(listRepository repository.TodoList, dbPool *pgxpool.Pool) *service {
	return &service{
		listRepository: listRepository,
		dbPool:         dbPool,
	}
}
