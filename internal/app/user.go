package app

import (
	"context"

	"github.com/nazip/grpc-auth/internal/api/user/v1/user"

	"github.com/nazip/grpc-auth/internal/repository"
	userRepository "github.com/nazip/grpc-auth/internal/repository/user"
	"github.com/nazip/grpc-auth/internal/service"
	userService "github.com/nazip/grpc-auth/internal/service/user"
)

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewServiceUser(
			s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}
	return s.userService
}

func (s *serviceProvider) UserApi(ctx context.Context) *user.Implementation {
	user.NewImplementation(s.UserService(ctx))
	if s.userApiImpl == nil {
		s.userApiImpl = user.NewImplementation(s.UserService(ctx))
	}
	return s.userApiImpl
}
