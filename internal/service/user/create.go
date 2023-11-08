package user

import (
	"context"
	model "github.com/nazip/grpc-auth/internal/models/service"
)

func (s *serviceUser) Create(ctx context.Context, u *model.User) (uint64, error) {
	var id uint64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.repository.Create(ctx, u)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
