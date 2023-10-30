package user

import (
	"context"
	model "github.com/nazip/grpc-auth/internal/models/service"
)

func (s *serviceUser) Create(ctx context.Context, u *model.User) (uint64, error) {
	id, err := s.repository.Create(ctx, u)
	if err != nil {
		return 0, err
	}

	return id, nil
}
