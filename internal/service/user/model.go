package user

import (
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
