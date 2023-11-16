package user

import (
	"context"

	model "github.com/nazip/grpc-auth/internal/models/service"
)

func (s *serviceUser) Update(ctx context.Context, u *model.User) error {
	return s.repository.Update(ctx, u)
}
