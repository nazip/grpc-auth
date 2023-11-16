package app

import (
	"context"

	authAPI "github.com/nazip/grpc-auth/internal/api/auth/v1/auth"
	"github.com/nazip/grpc-auth/internal/repository"
	authRepository "github.com/nazip/grpc-auth/internal/repository/auth"
	"github.com/nazip/grpc-auth/internal/service"
	authService "github.com/nazip/grpc-auth/internal/service/auth"
)

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepository.NewRepository(s.redisClient)
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewServiceAuth(s.AuthRepository(ctx), s.UserService(ctx))
	}
	return s.authService
}

func (s *serviceProvider) AuthApi(ctx context.Context) *authAPI.Implementation {
	if s.authApiImpl == nil {
		s.authApiImpl = authAPI.NewImplementation(s.AuthService(ctx))
	}
	return s.authApiImpl
}
