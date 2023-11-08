package user

import (
	"github.com/nazip/grpc-auth/internal/client/db"
	"github.com/nazip/grpc-auth/internal/repository"
	def "github.com/nazip/grpc-auth/internal/service"
)

var _ def.UserService = (*serviceUser)(nil)

type serviceUser struct {
	repository repository.UserRepository
	txManager  db.TxManager
}

func NewServiceUser(userRepository repository.UserRepository,
	txManager db.TxManager) def.UserService {
	return &serviceUser{
		repository: userRepository,
		txManager:  txManager,
	}
}
