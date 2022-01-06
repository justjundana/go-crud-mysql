package service

import (
	"crud-database/helper"
	"crud-database/models"
	"crud-database/repository"
	"time"
)

type UserService interface {
	GetUsersService() ([]models.User, error)
	GetUserService(id int) (models.User, error)
	CreateUserService(input helper.CreateUserRequest) (models.User, error)
	UpdateUserService(id int, input helper.EditUserRequest) (models.User, error)
	DeleteUserService(id int) (models.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) GetUsersService() ([]models.User, error) {
	users, err := s.repository.GetUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *userService) GetUserService(id int) (models.User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) CreateUserService(input helper.CreateUserRequest) (models.User, error) {
	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	createUser, err := s.repository.CreateUser(user)
	if err != nil {
		return createUser, err
	}
	return createUser, nil
}

func (s *userService) UpdateUserService(id int, input helper.EditUserRequest) (models.User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.UpdatedAt = time.Now()

	updateUser, err := s.repository.UpdateUser(user)
	if err != nil {
		return updateUser, err
	}
	return updateUser, nil
}

func (s *userService) DeleteUserService(id int) (models.User, error) {
	userID, err := s.GetUserService(id)
	if err != nil {
		return userID, err
	}

	deleteUser, err := s.repository.DeleteUser(userID)
	if err != nil {
		return deleteUser, err
	} else {
		return deleteUser, nil
	}
}
