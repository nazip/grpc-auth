package model

type Role int32

type CreateRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            Role
}

type CreateResponse struct {
	Id uint64
}
