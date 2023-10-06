package user

import (
	"context"
	"fmt"
	"time"

	desc "github.com/nazip/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Users map[uint64]User
type Users map[uint32]User

// User struct
type User struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      desc.Role `json:"role"`
	Password  string    `json:"password"`
	CreatedAt *timestamppb.Timestamp
	UpdatedAt *timestamppb.Timestamp
}

// NewUsers () Users
func NewUsers() Users {
	return make(map[uint32]User, 0)
}

// CreatedUser (ctx context.Context, user User) (User, error)
func (u Users) CreatedUser(_ context.Context, user User) (User, error) {
	// return exist user
	user1, ok := u[user.ID]
	if ok {
		return user1, nil
	}

	// return new user
	curDateTime := timestamppb.New(time.Now())
	temp := user
	temp.CreatedAt = curDateTime
	temp.UpdatedAt = curDateTime
	u[user.ID] = temp

	return u[user.ID], nil
}

// ReadUser (ctx context.Context, id uint64) (User, error)
func (u Users) ReadUser(_ context.Context, id uint32) (User, error) {
	user1, ok := u[id]
	if ok {
		return user1, nil
	}

	return User{}, fmt.Errorf("user %d not found", id)
}

// UpdateUser (ctx context.Context, user User) (User, error)
func (u Users) UpdateUser(_ context.Context, user User) (User, error) {
	foundedUser, ok := u[user.ID]
	if !ok {
		return User{}, fmt.Errorf("user %d not found", user.ID)
	}

	foundedUser.Email = user.Email
	foundedUser.Name = user.Name
	u[user.ID] = foundedUser

	return foundedUser, nil
}

// DeleteUser (ctx context.Context, userID uint64) error
func (u Users) DeleteUser(_ context.Context, userID uint32) error {
	_, ok := u[userID]
	if !ok {
		return fmt.Errorf("user %d not found", userID)
	}

	delete(u, userID)
	return nil
}
