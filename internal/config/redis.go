package config

import (
	"errors"
	"net"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

const (
	redisHostEnvName     = "REDIS_HOST"
	redisPortEnvName     = "REDIS_PORT"
	redisDatabaseEnvName = "REDIS_DB"
	redisUserEnvName     = "REDIS_USER"
	redisPasswordEnvName = "REDIS_PASSWORD"
)

type RedisConfig interface {
	Options() *redis.Options
}

type redisConfig struct {
	host     string
	port     string
	database string
	username string
	password string
}

func NewRedisConfig() (RedisConfig, error) {
	host := os.Getenv(redisHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("redis host not found")
	}

	port := os.Getenv(redisPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("redis port not found")
	}

	database := os.Getenv(redisDatabaseEnvName)
	if len(database) == 0 {
		return nil, errors.New("redis database not found")
	}

	username := os.Getenv(redisUserEnvName)
	if len(username) == 0 {
		return nil, errors.New("redis user name not found")
	}

	password := os.Getenv(redisPasswordEnvName)
	if len(username) == 0 {
		return nil, errors.New("redis password not found")
	}

	return &redisConfig{
		host:     host,
		port:     port,
		database: database,
		username: username,
		password: password,
	}, nil
}

func (cfg *redisConfig) Options() *redis.Options {
	db, _ := strconv.Atoi(cfg.database)
	return &redis.Options{
		Addr:     net.JoinHostPort(cfg.host, cfg.port),
		Username: cfg.username,
		Password: cfg.password,
		DB:       db,
	}
}
