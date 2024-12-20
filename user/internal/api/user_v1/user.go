package user_v1

import (
	"github.com/nazip/grpc-auth/internal/service"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
)

type UserServer struct {
	desc.UnimplementedUserV1Server
	Service service.Service
}
