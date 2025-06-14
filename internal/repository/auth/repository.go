package auth

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"github.com/ne4chelovek/base-rest-todo/internal/repository"
)

const (
	usersTable     = "users"
	idColum        = "id"
	nameColumn     = "name"
	userNameColumn = "username"
	passwordColumn = "password_hash"
)

type repo struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) repository.Authorization {
	return &repo{db: db}
}

func (r *repo) CreateUser(ctx context.Context, user *model.User) (int, error) {
	builderInsert := sq.Insert(usersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, userNameColumn, passwordColumn).
		Values(user.Name, user.Username, user.Password).
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

	return id, nil
}

func (r *repo) GetUser(ctx context.Context, username string) (*model.UserInfo, error) {
	builderSelect := sq.Select(idColum, userNameColumn, passwordColumn).
		From(usersTable).
		Where(sq.Eq{userNameColumn: username}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	user := &model.UserInfo{}

	err = r.db.QueryRow(ctx, query, args...).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
