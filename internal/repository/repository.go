package repository

import (
	"context"

	"github.com/nazip/grpc-auth/internal/models/repository"
	model "github.com/nazip/grpc-auth/internal/models/service"
)

type UserRepository interface {
	Create(ctx context.Context, req *model.User) (uint64, error)
	Get(ctx context.Context, id uint64) (*model.User, error)
	GetUser(ctx context.Context, username, password string) (*model.User, error)
	Update(ctx context.Context, req *model.User) error
	Delete(ctx context.Context, id uint64) error
}

type AuthRepository interface {
	Set(ctx context.Context, req *repository.Auth) error
	Get(ctx context.Context, id uint64) (*repository.Auth, error)
}

type AccessRepository interface {
	Check(ctx context.Context, id uint64) error
}
