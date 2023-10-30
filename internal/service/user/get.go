package user

import (
	"context"
	model "github.com/nazip/grpc-auth/internal/models/service"
)

func (s *serviceUser) Get(ctx context.Context, id uint64) (*model.User, error) {
	return s.repository.Get(ctx, id)
}
