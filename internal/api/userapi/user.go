package userapi

import (
	"context"
	conv "github.com/nazip/grpc-auth/internal/converter"
	"github.com/nazip/grpc-auth/internal/service"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserAPI struct {
	desc.UnimplementedUserV1Server
	serviceUser service.UserService
}

func NewUserAPI(rep service.UserService) *UserAPI {
	return &UserAPI{serviceUser: rep}
}

func (s *UserAPI) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.serviceUser.Create(ctx, conv.GetNewServiceUserForCreate(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *UserAPI) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.serviceUser.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: conv.GetProtoUserFromService(user),
	}, nil
}

func (s *UserAPI) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := s.serviceUser.Update(ctx, conv.GetNewServiceUserForUpdate(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *UserAPI) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.serviceUser.Delete(ctx, req.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
