package user_v1

import (
	"github.com/nazip/grpc-auth/internal/service"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
)

type User struct {
	desc.UnimplementedUserV1Server
	Service service.UserService
}

func NewUser(service service.UserService) *User {
	return &User{
		Service: service,
	}
}
