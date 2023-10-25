package converter

import (
	"github.com/nazip/grpc-auth/internal/model"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetUser(user model.User) *desc.GetResponse {
	var CreatedAt *timestamppb.Timestamp
	var UpdatedAt *timestamppb.Timestamp
	var email string
	if user.CreatedAt.Valid {
		CreatedAt = timestamppb.New(user.CreatedAt.Time)
	}
	if user.UpdatedAt.Valid {
		UpdatedAt = timestamppb.New(user.UpdatedAt.Time)
	}
	if user.Email.Valid {
		email = user.Email.String
	}

	return &desc.GetResponse{
		User: &desc.User{
			Id:        uint64(user.ID),
			Name:      user.Name,
			Email:     email,
			Role:      desc.Role(user.Role),
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		},
	}
}
