package repository

import (
	"context"
	"github.com/nazip/grpc-auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, request *model.CreateRequest) (*model.CreateResponse, error)
}
