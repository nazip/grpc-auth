package implamentation

import repo "github.com/nazip/grpc-auth/internal/repository"

type Service struct {
	repository repo.Repository
}

func NewService(repo repo.Repository) *Service {
	return &Service{
		repository: repo,
	}
}
