package app

import (
	"context"

	accessAPI "github.com/nazip/grpc-auth/internal/api/access/v1/access"
	"github.com/nazip/grpc-auth/internal/repository"
	accessRepository "github.com/nazip/grpc-auth/internal/repository/access"
	"github.com/nazip/grpc-auth/internal/service"
	accessService "github.com/nazip/grpc-auth/internal/service/access"
)

func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewRepository(s.redisClient)
	}

	return s.accessRepository
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewServiceAccess(s.AccessRepository(ctx))
	}
	return s.accessService
}

func (s *serviceProvider) AccessApi(ctx context.Context) *accessAPI.Implementation {
	if s.accessApiImpl == nil {
		s.accessApiImpl = accessAPI.NewImplementation(s.AccessService(ctx))
	}
	return s.accessApiImpl
}
