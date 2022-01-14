package service

import (
	"time"

	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/models"
	"github.com/justjundana/go-crud-mysql/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	LoginUserService(input helper.LoginUserRequest) (models.User, error)
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

func (s *userService) LoginUserService(input helper.LoginUserRequest) (models.User, error) {
	email := input.Email
	password := input.Password

	var user models.User
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err
}

func (s *userService) GetUsersService() ([]models.User, error) {
	users, err := s.repository.GetUsers()
	return users, err
}

func (s *userService) GetUserService(id int) (models.User, error) {
	user, err := s.repository.GetUser(id)
	return user, err
}

func (s *userService) CreateUserService(input helper.CreateUserRequest) (models.User, error) {
	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user.Password = string(passwordHash)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	createUser, err := s.repository.CreateUser(user)
	return createUser, err
}

func (s *userService) UpdateUserService(id int, input helper.EditUserRequest) (models.User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Email = input.Email
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user.Password = string(passwordHash)
	user.UpdatedAt = time.Now()

	updateUser, err := s.repository.UpdateUser(user)
	return updateUser, err
}

func (s *userService) DeleteUserService(id int) (models.User, error) {
	userID, err := s.GetUserService(id)
	if err != nil {
		return userID, err
	}

	deleteUser, err := s.repository.DeleteUser(userID)
	return deleteUser, err
}
