package access

import (
	"context"
	"github.com/nazip/grpc-auth/internal/repository"
	def "github.com/nazip/grpc-auth/internal/service"
)

type serviceAccess struct {
	repository repository.AccessRepository
}

func NewServiceAuth(accessRepository repository.AccessRepository) def.AccessService {
	return &serviceAccess{
		repository: accessRepository,
	}
}

func (s *serviceAccess) Check(ctx context.Context, userID uint64) error {
	if err := s.repository.Check(ctx, userID); err != nil {
		return err
	}

	return nil
}
