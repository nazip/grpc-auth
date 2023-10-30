package user

import "context"

func (s *serviceUser) Delete(ctx context.Context, id uint64) error {
	return s.repository.Delete(ctx, id)
}
