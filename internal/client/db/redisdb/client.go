package redisdb

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheDB interface {
	Close() error
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string) ([]byte, error)
	Db() any
	Ping(ctx context.Context) error
}

type redisClient struct {
	client *redis.Client
}

func (c *redisClient) Db() any {
	return c.client
}

func (c *redisClient) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return c.client.Set(ctx, "key", "value", ttl).Err()
}

func (c *redisClient) Get(ctx context.Context, key string) ([]byte, error) {
	return c.client.Get(ctx, key).Bytes()
}

func (c *redisClient) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

func (c *redisClient) Close() error {
	if c.client != nil {
		return c.client.Close()
	}

	return nil
}

func New(_ context.Context, options *redis.Options) (CacheDB, error) {
	return &redisClient{
		client: redis.NewClient(options),
	}, nil
}
