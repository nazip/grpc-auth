package implamentation

import (
	"context"
	"github.com/nazip/grpc-auth/internal/model"
)

func (s *Service) Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error) {
	return s.repository.Create(ctx, req)
}
