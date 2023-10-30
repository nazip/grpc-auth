package service

import (
	model "github.com/nazip/grpc-auth/internal/models/service"
	"golang.org/x/net/context"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) (uint64, error)
	Get(ctx context.Context, id uint64) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
}
