package v1

import (
	"context"
	desc "github.com/nazip/grpc-auth/pkg/auth_v1"
)

type API interface {
	Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error)
	RefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error)
	AccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error)
}
