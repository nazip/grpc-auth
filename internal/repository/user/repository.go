package user

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
)

const (
	tableName = "user"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repo {

	return &repo{db: db}
}

func (u *repo) Create(ctx context.Context, req *desc.CreateRequest) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn).
		Values(req.Name, req.Email, req.Password).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "note_repository.Create",
		QueryRaw: query,
	}

	return 0, nil
}

//func (u *User) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error)
//func (u *User) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error)
//func (u *User) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error)
