package auth

import (
	"context"
	"fmt"

	model "github.com/nazip/grpc-auth/internal/models/service"
	"github.com/nazip/grpc-auth/internal/repository"
	def "github.com/nazip/grpc-auth/internal/service"
)

type serviceAuth struct {
	repository  repository.AuthRepository
	userservice def.UserService
}

func NewServiceAuth(authRepository repository.AuthRepository, userservice def.UserService) def.AuthService {
	return &serviceAuth{
		repository:  authRepository,
		userservice: userservice,
	}
}

func (s *serviceAuth) Login(ctx context.Context, user *model.Auth) (string, error) {
	// check if user exists
	serviceUser, err := s.userservice.FindUser(ctx, user.UserName, user.Password)
	if err != nil {
		return "", err
	}

	// make jwt
	stab := "fake jwt for " + serviceUser.Name

	return stab, nil
}

func (s *serviceAuth) RefreshToken(ctx context.Context, token string) (string, error) {
	return "", fmt.Errorf("not implemented yet")
}

func (s *serviceAuth) AccessToken(ctx context.Context, token string) (string, error) {
	return "", fmt.Errorf("not implemented yet")
}
