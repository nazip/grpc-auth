package user

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nazip/grpc-auth/internal/repository"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type repo struct {
	q *Queries
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{New(db)}
}

func (u *repo) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	user, err := u.q.CreateUser(ctx, CreateUserParams{
		Name: req.Name,
		Email: sql.NullString{
			String: req.Email,
			Valid:  true,
		},
		Password: req.Password,
		Role:     int32(req.Role),
		CreatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: uint64(user.ID)}, nil
}

func (u *repo) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := u.q.GetUser(ctx, int64(req.GetId()))
	if err != nil {
		return nil, err
	}

	return GetUser(user), nil
}

func (u *repo) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	_, err := u.q.UpdateUser(ctx, UpdateUserParams{
		ID:   int64(req.Id),
		Name: req.Name,
		Email: sql.NullString{
			String: req.Email,
			Valid:  true,
		},
	})
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (u *repo) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := u.q.DeleteUser(ctx, int64(req.Id))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
