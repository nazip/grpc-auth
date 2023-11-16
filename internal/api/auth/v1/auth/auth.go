package auth

import (
	"context"

	converter "github.com/nazip/grpc-auth/internal/converter/auth/v1"
	"github.com/nazip/grpc-auth/internal/service"
	desc "github.com/nazip/grpc-auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	serviceAuth service.AuthService
}

func NewImplementation(service service.AuthService) *Implementation {
	return &Implementation{serviceAuth: service}
}

func (a *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {

	resp, err := a.serviceAuth.Login(ctx, converter.AuthServiceUser(req))
	if err != nil {
		return nil, err
	}

	return &desc.LoginResponse{
		RefreshToken: resp,
	}, nil
}

func (a *Implementation) RefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	return nil, nil
}
func (a *Implementation) AccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	return nil, nil
}
