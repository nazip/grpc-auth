package app

import (
	"context"
	userApiImpl "github.com/nazip/grpc-auth/internal/api/user_v1"
	"github.com/nazip/grpc-auth/internal/config"
	"github.com/nazip/grpc-auth/internal/config/env"
	"github.com/nazip/grpc-auth/internal/repository"
	repositoryImpl "github.com/nazip/grpc-auth/internal/repository/implamentation"
	"github.com/nazip/grpc-auth/internal/service"
	serviceImpl "github.com/nazip/grpc-auth/internal/service/implamentation"
	"log"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig
	service    service.UserService
	repository repository.UserRepository
	userApi    *userApiImpl.User
	//httpConfig   config.HTTPConfig
	//bankService  bank.Bank
	//dbClient     db.Client
	//loggerConfig config.LoggerConfig
	//metricConfig config.MetricConfig

}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig != nil {
		return s.grpcConfig
	}

	cfg, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %s", err.Error())
	}

	s.grpcConfig = cfg
	return s.grpcConfig
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.repository == nil {
		s.repository = repositoryImpl.NewRepository(ctx)
	}

	return s.repository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.service == nil {
		s.service = serviceImpl.NewService(s.UserRepository(ctx))
	}

	return s.service
}

func (s *serviceProvider) UserAPI(ctx context.Context) *userApiImpl.User {
	if s.userApi == nil {
		s.userApi = userApiImpl.NewUser(s.UserService(ctx))
	}

	return s.userApi
}
