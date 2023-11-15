package access

import (
	"context"
	"github.com/nazip/grpc-auth/internal/service"
	desc "github.com/nazip/grpc-auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server
	serviceAccess service.AccessService
}

func NewImplementation(service service.AccessService) *Implementation {
	return &Implementation{serviceAccess: service}
}

func (a *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
