package v1

import (
	"github.com/nazip/grpc-auth/internal/helpers"
	modelRepository "github.com/nazip/grpc-auth/internal/models/repository"
	modelService "github.com/nazip/grpc-auth/internal/models/service"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetProtoUserFromService(user *modelService.User) *desc.User {
	CreatedAt := timestamppb.Timestamp{
		Seconds: user.CreatedAt.Unix(),
		Nanos:   int32(user.CreatedAt.Nanosecond()),
	}

	UpdatedAt := timestamppb.Timestamp{
		Seconds: user.UpdatedAt.Unix(),
		Nanos:   int32(user.UpdatedAt.Nanosecond()),
	}

	return &desc.User{
		Id:        uint64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(user.Role),
		CreatedAt: &CreatedAt,
		UpdatedAt: &UpdatedAt,
	}
}

//func GetNewServiceUser(req *desc.CreateRequest) *modelRepository.User {
//	return &modelRepository.User{
//		Name:  req.Name,
//		Email: helpers.ToSqlString(req.Email),
//		Role:  int32(req.Role),
//	}
//}

func GetNewServiceUserForUpdate(req *desc.UpdateRequest) *modelService.User {
	return &modelService.User{
		ID:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	}
}

func GetNewServiceUserForCreate(req *desc.CreateRequest) *modelService.User {
	return &modelService.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     int32(req.Role),
	}
}

func ServiceUserFromRepo(repoUser *modelRepository.User) *modelService.User {

	return &modelService.User{
		ID:        repoUser.ID,
		Name:      repoUser.Name,
		Email:     helpers.FromSqlString(repoUser.Email),
		Password:  repoUser.Password,
		Role:      repoUser.Role,
		CreatedAt: helpers.FromSqlTime(repoUser.CreatedAt),
		UpdatedAt: helpers.FromSqlTime(repoUser.UpdatedAt),
	}
}

//func GetServiceUserFromProto(repoUser *modelRepository.User) *modelService.User {
//
//	return &modelService.User{
//		ID:        repoUser.ID,
//		Name:      repoUser.Name,
//		Email:     helpers.FromSqlString(repoUser.Email),
//		Password:  repoUser.Password,
//		Role:      repoUser.Role,
//		CreatedAt: helpers.FromSqlTime(repoUser.CreatedAt),
//		UpdatedAt: helpers.FromSqlTime(repoUser.UpdatedAt),
//	}
//}
