package service

import "time"

type User struct {
	ID        uint64
	Name      string
	Email     string
	Password  string
	Role      int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
