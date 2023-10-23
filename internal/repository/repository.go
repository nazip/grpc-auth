package repository

import (
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Repository interface {
	Create(ctx context.Context, req *desc.CreateRequest) (int64, error)
	Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error)
	Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error)
	Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error)
}
