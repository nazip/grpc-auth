package user_v1

import (
	"context"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
)

func (u UserServer) Create(context.Context, *desc.CreateRequest) (*desc.CreateResponse, error) {
	return &desc.CreateResponse{
		Id: 33,
	}, nil
}
