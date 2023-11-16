package tests

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	userAPI "github.com/nazip/grpc-auth/internal/api/user/v1/user"
	model "github.com/nazip/grpc-auth/internal/models/service"
	"github.com/nazip/grpc-auth/internal/service"
	serviceMocks "github.com/nazip/grpc-auth/internal/service/mocks"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUserAPI_Get(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		id    = gofakeit.Uint64()
		name  = gofakeit.Name()
		email = gofakeit.Email()

		req = &desc.GetRequest{
			Id: id,
		}

		user = &model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Role:      0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		resp = &desc.GetResponse{
			User: &desc.User{
				Id:    id,
				Name:  name,
				Email: email,
				Role:  0,
				CreatedAt: &timestamppb.Timestamp{
					Seconds: user.CreatedAt.Unix(),
					Nanos:   int32(user.CreatedAt.Nanosecond()),
				},
				UpdatedAt: &timestamppb.Timestamp{
					Seconds: user.UpdatedAt.Unix(),
					Nanos:   int32(user.UpdatedAt.Nanosecond()),
				},
			},
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: resp,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(user, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := userAPI.NewImplementation(userServiceMock)

			resp, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want.User, resp.User)
		})
	}

}
