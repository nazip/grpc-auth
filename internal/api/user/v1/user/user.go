package user

import (
	"context"

	conv "github.com/nazip/grpc-auth/internal/converter/user/v1"
	"github.com/nazip/grpc-auth/internal/service"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	serviceUser service.UserService
}

func NewImplementation(srv service.UserService) *Implementation {
	return &Implementation{serviceUser: srv}
}

func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.serviceUser.Create(ctx, conv.GetNewServiceUserForCreate(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.serviceUser.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: conv.GetProtoUserFromService(user),
	}, nil
}

func (s *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := s.serviceUser.Update(ctx, conv.GetNewServiceUserForUpdate(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.serviceUser.Delete(ctx, req.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
