package implamentation

import (
	"context"
	"github.com/nazip/grpc-auth/internal/model"
)

func (r *Repository) Create(ctx context.Context, request *model.CreateRequest) (*model.CreateResponse, error) {
	return &model.CreateResponse{
		Id: 444,
	}, nil
}
