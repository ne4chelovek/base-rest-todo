package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
)

type QueryRunner interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type Authorization interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
	GetUser(ctx context.Context, username string) (*model.UserInfo, error)
}

type TodoList interface {
	WithTx(tx pgx.Tx) TodoList
	Create(ctx context.Context, userId int, list *model.TodoList) (int, error)
	GetAll(ctx context.Context, userId int) ([]*model.TodoList, error)
	GetById(ctx context.Context, userId int, listId int) (*model.TodoList, error)
	Update(ctx context.Context, userId int, listId int, input *model.UpdateListInput) error
	Delete(ctx context.Context, userId int, listId int) error
}

type TodoItem interface {
	WithTx(tx pgx.Tx) TodoItem
	Create(ctx context.Context, listId int, item *model.TodoItem) (int, error)
	GetAllItem(ctx context.Context, userId, listId int) ([]*model.TodoItem, error)
	GetById(ctx context.Context, userId, itemId int) (*model.TodoItem, error)
	Delete(ctx context.Context, userId, itemId int) error
	Update(ctx context.Context, userId, itemId int, input *model.UpdateItemInput) error
}
