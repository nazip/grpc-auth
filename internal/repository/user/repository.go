package user

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/nazip/grpc-auth/internal/client/db"
	"github.com/nazip/grpc-auth/internal/converter/user/v1"
	"github.com/nazip/grpc-auth/internal/helpers"
	modelRepository "github.com/nazip/grpc-auth/internal/models/repository"
	modelService "github.com/nazip/grpc-auth/internal/models/service"

	"github.com/nazip/grpc-auth/internal/repository"
	"time"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db}
}

func (u *repo) Create(ctx context.Context, req *modelService.User) (uint64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn, createdAtColumn).
		Values(req.Name, req.Email, req.Password, req.Role,
			sql.NullTime{Time: time.Now(), Valid: true}).
		Suffix("RETURNING id")
	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id uint64
	err = u.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *repo) Get(ctx context.Context, id uint64) (*modelService.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn,
		passwordColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	user := modelRepository.User{}
	err = u.db.DB().QueryRowContext(ctx, q, args...).Scan(&user.ID, &user.Name,
		&user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return v1.ServiceUserFromRepo(&user), nil
}

func (u *repo) Update(ctx context.Context, req *modelService.User) error {
	builder := sq.Update(tableName).
		Set(nameColumn, req.Name).
		Set(emailColumn, req.Email).
		Set(updatedAtColumn, helpers.ToSqlTime(time.Now().UTC())).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: req.ID})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = u.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (u *repo) Delete(ctx context.Context, id uint64) error {
	builder := sq.Delete(tableName).Where(sq.Eq{idColumn: id})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = u.db.DB().ExecContext(ctx, q, args)
	if err != nil {
		return err
	}

	return nil
}
