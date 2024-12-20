package repository

import (
	"context"
	"github.com/nazip/grpc-auth/internal/model"
)

type Repository interface {
	Create(ctx context.Context, request model.CreateRequest) model.CreateResponse
}
