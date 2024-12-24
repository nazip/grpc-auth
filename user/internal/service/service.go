package service

import (
	"context"
	"github.com/nazip/grpc-auth/internal/model"
)

type UserService interface {
	Create(context.Context, *model.CreateRequest) (*model.CreateResponse, error)
}
