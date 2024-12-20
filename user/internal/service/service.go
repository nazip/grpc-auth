package service

import (
	"context"
	"github.com/nazip/grpc-auth/internal/model"
)

type Service interface {
	Create(context.Context, model.CreateRequest) (model.CreateResponse, error)
}
