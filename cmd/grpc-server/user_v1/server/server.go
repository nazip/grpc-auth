package server

import (
	"context"
	"github.com/nazip/grpc-auth/internal/repository"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedUserV1Server
	repositoryUser repository.NoteRepository
}

// NewServer returns *server
func NewServer(rep repository.NoteRepository) *server {
	return &server{repositoryUser: rep}
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	return s.repositoryUser.Create(ctx, req)
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	return s.repositoryUser.Get(ctx, req)
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	return s.repositoryUser.Update(ctx, req)
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	return s.repositoryUser.Delete(ctx, req)
}
