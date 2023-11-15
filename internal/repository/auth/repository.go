package auth

import (
	"context"
	"encoding/json"
	"github.com/nazip/grpc-auth/internal/client/db"
	model "github.com/nazip/grpc-auth/internal/models/repository"
	"github.com/nazip/grpc-auth/internal/repository"
	"strconv"
)

type repo struct {
	db db.CacheDB
}

func NewRepository(db db.CacheDB) repository.AuthRepository {
	return &repo{db}
}

func (u *repo) Set(ctx context.Context, req *model.Auth) error {
	rez, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return u.db.Set(ctx, strconv.Itoa(int(req.ID)), rez, req.TTL)
}

func (u *repo) Get(ctx context.Context, id uint64) (*model.Auth, error) {
	rez, err := u.db.Get(ctx, strconv.Itoa(int(id)))
	if err != nil {
		return nil, err
	}

	auth := &model.Auth{}
	err = json.Unmarshal(rez, auth)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
