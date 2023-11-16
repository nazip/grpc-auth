package app

import (
	"context"
	"log"

	accessAPI "github.com/nazip/grpc-auth/internal/api/access/v1/access"
	authAPI "github.com/nazip/grpc-auth/internal/api/auth/v1/auth"
	userAPI "github.com/nazip/grpc-auth/internal/api/user/v1/user"
	"github.com/nazip/grpc-auth/internal/client/db"
	"github.com/nazip/grpc-auth/internal/client/db/pg"
	"github.com/nazip/grpc-auth/internal/client/db/redisdb"
	"github.com/nazip/grpc-auth/internal/client/db/transaction"
	"github.com/nazip/grpc-auth/internal/closer"
	"github.com/nazip/grpc-auth/internal/config"
	"github.com/nazip/grpc-auth/internal/repository"
	"github.com/nazip/grpc-auth/internal/service"
)

type serviceProvider struct {
	pgConfig    config.PGConfig
	dbClient    db.Client
	redisClient redisdb.CacheDB
	txManager   db.TxManager

	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
	redisConfig   config.RedisConfig

	userRepository repository.UserRepository
	userService    service.UserService
	userApiImpl    *userAPI.Implementation

	authRepository repository.AuthRepository
	authService    service.AuthService
	authApiImpl    *authAPI.Implementation

	accessRepository repository.AccessRepository
	accessService    service.AccessService
	accessApiImpl    *accessAPI.Implementation
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

func (s *serviceProvider) RedisClient(ctx context.Context) redisdb.CacheDB {
	if s.redisClient == nil {
		//cl := redis.NewClient(s.RedisConfig().Options())
		redisClient, err := redisdb.New(ctx, s.RedisConfig().Options())
		if err != nil {
			log.Fatalf("failed to create redis client: %v", err)
		}

		if err := redisClient.Ping(ctx); err != nil {
			log.Fatalf("failed to create redis client: %v", err)
		}
		s.redisClient = redisClient
	}

	return s.redisClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}
