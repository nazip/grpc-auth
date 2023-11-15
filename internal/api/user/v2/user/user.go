package user

import (
	"context"
	conv "github.com/nazip/grpc-auth/internal/converter/user/v2"
	"github.com/nazip/grpc-auth/internal/service"
	desc "github.com/nazip/grpc-auth/pkg/user_v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type API struct {
	desc.UnimplementedUserV2Server
	serviceUser service.UserService
}

func NewAPI(srv service.UserService) *API {
	return &API{serviceUser: srv}
}

func (s *API) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.serviceUser.Create(ctx, conv.GetNewServiceUserForCreate(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *API) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.serviceUser.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: conv.GetProtoUserFromService(user),
	}, nil
}

func (s *API) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := s.serviceUser.Update(ctx, conv.GetNewServiceUserForUpdate(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *API) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.serviceUser.Delete(ctx, req.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
