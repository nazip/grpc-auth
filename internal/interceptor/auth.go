package interceptor

import (
	"context"
	modelService "github.com/nazip/grpc-auth/internal/models/service"
	"google.golang.org/grpc"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if _, ok := req.(modelService.User); ok {
		// check user is authenticated
		return handler(ctx, req)
	}

	if _, ok := req.(uint64); ok {
		// check user is authenticated
		return handler(ctx, req)
	}

	return handler(ctx, req)
}
