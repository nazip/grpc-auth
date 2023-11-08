package app

import (
	"context"
	"github.com/nazip/grpc-auth/internal/api/userapi"
	"github.com/nazip/grpc-auth/internal/client/db"
	"github.com/nazip/grpc-auth/internal/client/db/pg"
	"github.com/nazip/grpc-auth/internal/client/db/transaction"
	"github.com/nazip/grpc-auth/internal/closer"
	"github.com/nazip/grpc-auth/internal/config"
	"github.com/nazip/grpc-auth/internal/repository"
	userRepository "github.com/nazip/grpc-auth/internal/repository/user"
	"github.com/nazip/grpc-auth/internal/service"
	userService "github.com/nazip/grpc-auth/internal/service/user"
	"log"
)

type serviceProvider struct {
	pgConfig       config.PGConfig
	dbClient       db.Client
	txManager      db.TxManager
	grpcConfig     config.GRPCConfig
	userRepository repository.UserRepository
	userService    service.UserService
	userApi        *userapi.UserAPI
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

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

func (s *serviceProvider) UserApi(ctx context.Context) *userapi.UserAPI {
	if s.userApi == nil {
		s.userApi = userapi.NewUserAPI(s.UserService(ctx))
	}
	return s.userApi
}
