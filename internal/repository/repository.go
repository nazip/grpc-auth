package repository

import (
	"context"
	model "github.com/nazip/grpc-auth/internal/service/user"
)

type UserRepository interface {
	Create(ctx context.Context, req *model.User) (uint64, error)
	Get(ctx context.Context, id uint64) (*model.User, error)
	Update(ctx context.Context, req *model.User) error
	Delete(ctx context.Context, id uint64) error
}
