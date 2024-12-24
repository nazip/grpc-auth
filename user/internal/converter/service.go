package converter

import (
	"github.com/nazip/grpc-auth/internal/model"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
)

func FromProtoToService(request *desc.CreateRequest) *model.CreateRequest {
	return &model.CreateRequest{
		Name:            request.Name,
		Email:           request.Email,
		Password:        request.Password,
		PasswordConfirm: request.PasswordConfirm,
		Role:            model.Role(request.Role),
	}

}

func FromServiceToProto(request model.CreateResponse) desc.CreateResponse {
	return desc.CreateResponse{
		Id: request.Id,
	}
}
