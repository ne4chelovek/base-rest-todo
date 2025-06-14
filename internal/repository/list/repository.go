package list

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"github.com/ne4chelovek/base-rest-todo/internal/repository"
	"golang.org/x/net/context"
)

const (
	todoListsTable = "todo_lists"
	usersListTable = "users_lists"

	todoListID          = "id"
	todoListTitle       = "title"
	todoListDescription = "description"

	usersListUserId = "user_id"
	usersListListId = "list_id"
)

type repo struct {
	db repository.QueryRunner
}

func NewListRepository(db *pgxpool.Pool) repository.TodoList {
	return &repo{db: db}
}

func (r *repo) WithTx(tx pgx.Tx) repository.TodoList {
	return &repo{db: tx}
}

func (r *repo) Create(ctx context.Context, userId int, list *model.TodoList) (int, error) {
	builderInsert := sq.Insert(todoListsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(todoListTitle, todoListDescription).
		Values(list.Title, list.Description).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	var id int
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	builderListInsert := sq.Insert(usersListTable).
		PlaceholderFormat(sq.Dollar).
		Columns(usersListUserId, usersListListId).
		Values(userId, id)

	query, args, err = builderListInsert.ToSql()
	if err != nil {
		return 0, err
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return 0, fmt.Errorf("failed to create user-list association")
	}

	return id, nil
}

func (r *repo) GetAll(ctx context.Context, userId int) ([]*model.TodoList, error) {
	query, args, err := sq.Select("tl."+todoListID, "tl."+todoListTitle, "tl."+todoListDescription).
		From(todoListsTable + " tl").
		Join(usersListTable + " ul ON tl.id = ul.list_id").
		Where(sq.Eq{"ul." + usersListUserId: userId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	lists := []*model.TodoList{}
	for rows.Next() {
		list := &model.TodoList{}
		if err := rows.Scan(&list.Id, &list.Title, &list.Description); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		lists = append(lists, list)
	}
	return lists, nil
}

func (r *repo) GetById(ctx context.Context, userId int, listId int) (*model.TodoList, error) {
	query, args, err := sq.Select(
		"tl."+todoListID, "tl."+todoListTitle, "tl."+todoListDescription,
	).
		From(todoListsTable + " tl").
		Join(usersListTable + " ul ON tl.id = ul.list_id").
		Where(sq.Eq{
			"ul." + usersListUserId: userId,
			"ul." + usersListListId: listId,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	list := &model.TodoList{}

	err = r.db.QueryRow(ctx, query, args...).Scan(&list.Id, &list.Title, &list.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to get list: %v", err)
	}

	return list, nil
}

func (r *repo) Update(ctx context.Context, userId int, listId int, input *model.UpdateListInput) error {
	builderUpdate := sq.Update(todoListsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Expr(fmt.Sprintf("id IN (SELECT list_id FROM %s WHERE list_id = ? AND user_id = ?)", usersListTable), listId, userId))

	if input.Title != nil {
		builderUpdate = builderUpdate.Set(todoListTitle, *input.Title)
	}
	if input.Description != nil {
		builderUpdate = builderUpdate.Set(todoListDescription, *input.Description)
	}

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("failed to update list - not found or no permission")
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, userId int, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListsTable, usersListTable)

	_, err := r.db.Exec(ctx, query, userId, listId)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}
