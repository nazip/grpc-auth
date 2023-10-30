package user

import (
	"github.com/nazip/grpc-auth/internal/repository"
	def "github.com/nazip/grpc-auth/internal/service"
)

var _ def.UserService = (*serviceUser)(nil)

type serviceUser struct {
	repository repository.UserRepository
}

func NewServiceUser(userRepository repository.UserRepository) *serviceUser {
	return &serviceUser{
		repository: userRepository,
	}
}
