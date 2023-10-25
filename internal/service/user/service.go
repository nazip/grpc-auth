package user

import "github.com/nazip/grpc-auth/internal/repository"

type serv struct {
	repository repository.UserRepository
}
