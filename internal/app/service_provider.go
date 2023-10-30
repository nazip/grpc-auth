package app

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nazip/grpc-auth/internal/api/userapi"
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
	pgxpool        *pgxpool.Pool
	grpcConfig     config.GRPCConfig
	userRepository repository.UserRepository
	userService    service.UserService
	userImpl       *userapi.UserAPI
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PgxPool(ctx context.Context) *pgxpool.Pool {
	if s.pgxpool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("error connecting to pgxpool: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("error connecting to pgxpool: %v", err)
		}
		s.pgxpool = pool

		closer.Add(func() error {
			s.pgxpool.Close()
			return nil
		})
	}
	return s.pgxpool
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
		s.userRepository = userRepository.NewRepository(s.PgxPool(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewServiceUser(s.UserRepository(ctx))
	}
	return s.userService
}
