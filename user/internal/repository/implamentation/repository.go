package implamentation

import (
	"context"
	desc "github.com/nazip/grpc-auth/internal/repository"
)

var _ desc.UserRepository = &Repository{}

type Repository struct {
}

func NewRepository(_ context.Context) *Repository {
	return &Repository{}
}
