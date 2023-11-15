package v1

import (
	repomodel "github.com/nazip/grpc-auth/internal/models/repository"
	servicemodel "github.com/nazip/grpc-auth/internal/models/service"
)

func AuthUser(repoAuthUser *repomodel.Auth) *servicemodel.Auth {
	return &servicemodel.Auth{
		ID:    repoAuthUser.ID,
		Token: repoAuthUser.Token,
		TTL:   repoAuthUser.TTL,
	}
}
