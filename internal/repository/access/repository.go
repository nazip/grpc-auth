package access

import (
	"context"
	"fmt"

	"github.com/nazip/grpc-auth/internal/client/db"
	def "github.com/nazip/grpc-auth/internal/repository"
)

type repo struct {
	db db.CacheDB
}

func NewRepository(db db.CacheDB) def.AccessRepository {
	return &repo{db}
}

func (r *repo) Check(ctx context.Context, userID uint64) error {
	return fmt.Errorf("not implemented yet")
}
