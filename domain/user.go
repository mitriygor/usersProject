package domain

import (
	"github.com/mitriygor/usersProject/dto"
	"github.com/mitriygor/usersProjectLib/errors"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (c User) ToDto() dto.UserResponse {
	return dto.UserResponse{
		Email:     c.Email,
		FirstName: c.FirstName,
		LastName:  c.LastName,
	}
}

type UserRepository interface {
	FindAll(status string) ([]User, *errors.AppError)
}
