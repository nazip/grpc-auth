package v1

import (
	"context"
	desc "github.com/nazip/grpc-auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type API interface {
	Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error)
}
