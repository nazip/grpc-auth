package tests

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	userAPI "github.com/nazip/grpc-auth/internal/api/user/v2/user"
	model "github.com/nazip/grpc-auth/internal/models/service"
	"github.com/nazip/grpc-auth/internal/service"
	serviceMocks "github.com/nazip/grpc-auth/internal/service/mocks"
	desc "github.com/nazip/grpc-auth/pkg/user_v2"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestUserAPI_Update(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		id    = gofakeit.Uint64()
		name  = gofakeit.Name()
		email = gofakeit.Email()

		req = &desc.UpdateRequest{
			Id:    id,
			Name:  name,
			Email: email,
		}

		user = &model.User{
			ID:    id,
			Name:  name,
			Email: email,
			Role:  0,
		}

		resp = &emptypb.Empty{}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
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
				mock.UpdateMock.Expect(ctx, user).Return(nil)
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

			resp, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resp)
		})
	}

}
