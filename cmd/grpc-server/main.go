package main

import (
	"context"
	"flag"
	"github.com/nazip/grpc-auth/internal/repository/user"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nazip/grpc-auth/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/nazip/grpc-auth/cmd/grpc-server/user_v1/server"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "local.env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// grpc config
	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	// pg config
	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)

	repUser := user.NewRepository(pool)
	srv := server.NewServer(repUser)
	desc.RegisterUserV1Server(s, srv)

	lis, err := net.Listen("tcp", grpcConfig.GRPCAddress())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-stop
	s.GracefulStop()
}
