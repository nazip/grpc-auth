package server

import (
	"context"

	"github.com/brianvoe/gofakeit"

	"github.com/nazip/grpc-auth/cmd/grpc-server/user_v1/user"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedUserV1Server
	users user.Users
}

// NewServer returns *server
func NewServer() *server {
	return &server{users: user.NewUsers()}
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id := gofakeit.Uint64()
	return &desc.CreateResponse{
		Id: id,
	}, nil

	// if req.Password != req.PasswordConfirm {
	// 	return &desc.CreateResponse{}, fmt.Errorf("password != password confirm")
	// }

	// user, err := s.users.CreatedUser(ctx, user.User{
	// 	ID:        id,
	// 	Name:      req.Name,
	// 	Email:     req.Email,
	// 	Password:  req.Password,
	// 	Role:      req.Role,
	// 	CreatedAt: timestamppb.New(time.Now()),
	// 	UpdatedAt: timestamppb.New(time.Now()),
	// })
	// if err != nil {
	// 	return &desc.CreateResponse{}, err
	// }

	// return &desc.CreateResponse{
	// 	Id: user.ID,
	// }, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.users.ReadUser(ctx, req.Id)
	if err != nil {
		return &desc.GetResponse{}, err
	}

	return &desc.GetResponse{
		User: &desc.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	_, err := s.users.UpdateUser(ctx, user.User{
		ID:    uint64(req.Id),
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.users.DeleteUser(ctx, uint64(req.Id))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
