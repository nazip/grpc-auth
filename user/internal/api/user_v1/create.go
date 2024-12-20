package user_v1

import (
	"context"
	"github.com/nazip/grpc-auth/internal/converter"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
)

func (u UserServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	resp, err := u.Service.Create(ctx, converter.FromProtoToService(*req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: 33,
	}, nil
}
