package user

import (
	"context"

	model "github.com/nazip/grpc-auth/internal/models/service"
)

func (s *serviceUser) FindUser(ctx context.Context, username, password string) (*model.User, error) {
	return s.repository.GetUser(ctx, username, password)
}
