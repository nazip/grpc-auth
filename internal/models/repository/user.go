package repository

import "database/sql"

type User struct {
	ID        uint64
	Name      string
	Email     sql.NullString
	Password  string
	Role      int32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
