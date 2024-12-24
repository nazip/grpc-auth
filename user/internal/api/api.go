package api

import (
	"context"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
)

type UserAPI interface {
	Create(context.Context, *desc.CreateRequest) (*desc.CreateResponse, error)
}
