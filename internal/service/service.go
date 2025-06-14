package service

import (
	"context"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
)

type Authorization interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
	GetUser(ctx context.Context, username string) (*model.UserInfo, error)
}

type TodoList interface {
	Create(ctx context.Context, userId int, list *model.TodoList) (int, error)
	GetAll(ctx context.Context, userId int) ([]*model.TodoList, error)
	GetById(ctx context.Context, userId int, listId int) (*model.TodoList, error)
	Update(ctx context.Context, userId int, listId int, input *model.UpdateListInput) error
	Delete(ctx context.Context, userId int, listId int) error
}

type TodoItem interface {
	Create(ctx context.Context, userId, listId int, item *model.TodoItem) (int, error)
	GetAllItem(ctx context.Context, userId, listId int) ([]*model.TodoItem, error)
	GetById(ctx context.Context, userId, itemId int) (*model.TodoItem, error)
	Delete(ctx context.Context, userId, itemId int) error
	Update(ctx context.Context, userId, itemId int, input *model.UpdateItemInput) error
}

type Token interface {
	GenerateToken(ctx context.Context, username string, password string) (string, error)
	ParseToken(token string) (int, error)
}
