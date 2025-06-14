package item

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
	todoItemsTable = "todo_items"
	listsItemTable = "lists_items"
	usersListTable = "users_lists"

	todoListID          = "id"
	todoListTitle       = "title"
	todoListDescription = "description"

	ItemId = "item_id"
	ListId = "list_id"

	todoItemDoneColumn = "done"
)

type repo struct {
	db repository.QueryRunner
}

func NewItemRepository(db *pgxpool.Pool) repository.TodoItem {
	return &repo{db: db}
}

func (r *repo) WithTx(tx pgx.Tx) repository.TodoItem {
	return &repo{db: tx}
}

func (r *repo) Create(ctx context.Context, listId int, item *model.TodoItem) (int, error) {
	builderInsert := sq.Insert(todoItemsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(todoListTitle, todoListDescription).
		Values(item.Title, item.Description).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	var itemId int
	err = r.db.QueryRow(ctx, query, args...).Scan(&itemId)
	if err != nil {
		return 0, err
	}

	builderInsertItem := sq.Insert(listsItemTable).
		PlaceholderFormat(sq.Dollar).
		Columns(ListId, ItemId).
		Values(listId, itemId)

	query, args, err = builderInsertItem.ToSql()
	if err != nil {
		return 0, err
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return 0, fmt.Errorf("no item found with id %d", listId)
	}

	return itemId, nil
}

func (r *repo) GetAllItem(ctx context.Context, userId, listId int) ([]*model.TodoItem, error) {
	builderSelect := sq.Select(
		"ti.id",
		"ti.title",
		"ti.description",
		"ti.done").
		PlaceholderFormat(sq.Dollar).
		From(todoItemsTable + " ti").
		Join(listsItemTable + " li ON li.item_id = ti.id").
		Join(usersListTable + " ul ON ul.list_id = li.list_id").
		Where(sq.And{
			sq.Eq{"li.list_id": listId},
			sq.Eq{"ul.user_id": userId},
		})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	items := []*model.TodoItem{}
	for rows.Next() {
		item := &model.TodoItem{}
		err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return items, nil
}

func (r *repo) GetById(ctx context.Context, userId, itemId int) (*model.TodoItem, error) {
	builderSelect := sq.Select(
		"ti.id",
		"ti.title",
		"ti.description",
		"ti.done").
		PlaceholderFormat(sq.Dollar).
		From(todoItemsTable + " ti").
		Join(listsItemTable + " li ON li.item_id = ti.id").
		Join(usersListTable + " ul ON ul.list_id = li.list_id").
		Where(sq.And{
			sq.Eq{"ti.id": itemId},
			sq.Eq{"ul.user_id": userId},
		})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	item := &model.TodoItem{}
	err = r.db.QueryRow(ctx, query, args...).Scan(&item.Id, &item.Title, &item.Description, &item.Done)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return item, nil
}

func (r *repo) Delete(ctx context.Context, userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listsItemTable, usersListTable)
	result, err := r.db.Exec(ctx, query, userId, itemId)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no item found with id %d", itemId)
	}

	return nil
}

func (r *repo) Update(ctx context.Context, userId, itemId int, input *model.UpdateItemInput) error {
	builderUpdate := sq.Update(todoItemsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Expr(fmt.Sprintf(`
            id IN (
                SELECT li.item_id 
                FROM %s li 
                JOIN %s ul ON li.list_id = ul.list_id 
                WHERE ul.user_id = ? AND li.item_id = ?
            )`, listsItemTable, usersListTable), userId, itemId))

	if input.Title != nil {
		builderUpdate = builderUpdate.Set("title", *input.Title)
	}
	if input.Description != nil {
		builderUpdate = builderUpdate.Set("description", *input.Description)
	}
	if input.Done != nil {
		builderUpdate = builderUpdate.Set("done", *input.Done)
	}

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build update query: %w", err)
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute update: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no item found with id %d", itemId)
	}

	return nil
}
