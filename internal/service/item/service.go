package item

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ne4chelovek/base-rest-todo/internal/repository"
)

type service struct {
	itemRepository repository.TodoItem
	listRepository repository.TodoList
	dbPool         *pgxpool.Pool
}

func NewService(itemRepository repository.TodoItem, listRepository repository.TodoList, dbPool *pgxpool.Pool) *service {
	return &service{
		itemRepository: itemRepository,
		listRepository: listRepository,
		dbPool:         dbPool,
	}
}
