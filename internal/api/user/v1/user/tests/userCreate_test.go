package tests

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	userAPI "github.com/nazip/grpc-auth/internal/api/user/v1/user"
	model "github.com/nazip/grpc-auth/internal/models/service"
	"github.com/nazip/grpc-auth/internal/service"
	serviceMocks "github.com/nazip/grpc-auth/internal/service/mocks"
	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAPI_Create(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Animal()
		id       = gofakeit.Uint64()

		req = &desc.CreateRequest{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     0,
		}

		resp = &desc.CreateResponse{
			Id: id,
		}

		user = &model.User{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     0,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mock.CreateMock.Expect(ctx, user).Return(id, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := userAPI.NewAPI(userServiceMock)

			resp, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want.Id, resp.Id)
		})
	}

}
