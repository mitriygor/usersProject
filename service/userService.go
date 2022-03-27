package service

import (
	"github.com/mitriygor/usersProject/domain"
	"github.com/mitriygor/usersProject/dto"
	"github.com/mitriygor/usersProjectLib/errors"
)

type UserService interface {
	GetAllUsers(string) ([]dto.UserResponse, *errors.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) GetAllUsers(status string) ([]dto.UserResponse, *errors.AppError) {
	users, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.UserResponse, 0)
	for _, c := range users {
		response = append(response, c.ToDto())
	}
	return response, err
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
