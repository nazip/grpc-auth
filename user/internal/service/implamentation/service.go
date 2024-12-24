package implamentation

import (
	repo "github.com/nazip/grpc-auth/internal/repository"
	desc "github.com/nazip/grpc-auth/internal/service"
)

var _ desc.UserService = &Service{}

type Service struct {
	repository repo.UserRepository
}

func NewService(repo repo.UserRepository) *Service {
	return &Service{
		repository: repo,
	}
}
