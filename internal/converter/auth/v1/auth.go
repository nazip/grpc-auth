package v1

import (
	"github.com/nazip/grpc-auth/internal/models/service"
	desc "github.com/nazip/grpc-auth/pkg/auth_v1"
)

func AuthServiceUser(req *desc.LoginRequest) *service.Auth {
	return &service.Auth{
		UserName: req.Username,
		Password: req.Password,
	}
}
