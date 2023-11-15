package repository

import "time"

type Auth struct {
	ID    uint64
	Token string
	TTL   time.Duration
}
