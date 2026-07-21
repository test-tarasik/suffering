package models

import "github.com/google/uuid"

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type RegisterResponse struct {
	ID string `json:"id"`
}

func NewUser(email string) User {
	return User{
		ID:    uuid.New().String(),
		Email: email,
	}
}
