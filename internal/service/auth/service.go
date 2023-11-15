package auth

import (
	"context"
	"fmt"
	model "github.com/nazip/grpc-auth/internal/models/service"
	"github.com/nazip/grpc-auth/internal/repository"
	def "github.com/nazip/grpc-auth/internal/service"
)

type serviceAuth struct {
	repository repository.AuthRepository
}

func NewServiceAuth(authRepository repository.AuthRepository) def.AuthService {
	return &serviceAuth{
		repository: authRepository,
	}
}

func (s *serviceAuth) Login(ctx context.Context, user *model.Auth) error {
	return fmt.Errorf("not implemented yet")
}

func (s *serviceAuth) RefreshToken(ctx context.Context, token string) (string, error) {
	return "", fmt.Errorf("not implemented yet")
}

func (s *serviceAuth) AccessToken(ctx context.Context, token string) (string, error) {
	return "", fmt.Errorf("not implemented yet")
}
