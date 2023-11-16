package v1

import (
	"context"

	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type API interface {
	Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error)
	Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error)
	Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error)
	Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error)
}
