package service

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/helpers"
	"backend-inventory-app/repository"
	userPayload "backend-inventory-app/web/users"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateNewUser(request userPayload.UserCreateRequest) (userPayload.UserResponse, error)
	LoginUser(request userPayload.UserLoginRequest) (userPayload.UserResponse, error)
	GetUserByID(userID int) (userPayload.UserResponse, error)
	UpdateUser(userID int, request userPayload.UserUpdateRequest) (userPayload.UserResponse, error)
	DeleteUser(userID int) (bool, error)
	GetUsers() ([]userPayload.UserResponse, error)
}

type newUserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &newUserService{repository}
}

func (s *newUserService) CreateNewUser(request userPayload.UserCreateRequest) (userPayload.UserResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return userPayload.UserResponse{}, err
	}

	data := entity.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: string(passwordHash),
	}

	user, err := s.repository.Create(data)
	if err != nil {
		return userPayload.UserResponse{}, err
	}
	token, err := helpers.GenerateToken(user.ID)

	return helpers.ToUserResponse(user, token), nil
}

func (s *newUserService) LoginUser(request userPayload.UserLoginRequest) (userPayload.UserResponse, error) {
	// Cek Email
	user, err := s.repository.FindByEmail(request.Email)
	if err != nil || user.ID == 0 {
		return userPayload.UserResponse{}, errors.New("email not found")
	}

	// Cek Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return userPayload.UserResponse{}, errors.New("password wrong")
	}

	token, err := helpers.GenerateToken(user.ID)

	return helpers.ToUserResponse(user, token), nil
}

func (s *newUserService) GetUserByID(userID int) (userPayload.UserResponse, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil || user.ID == 0 {
		return helpers.ToUserResponse(user, ""), errors.New("not user found")
	}

	return helpers.ToUserResponse(user, ""), nil
}

func (s *newUserService) UpdateUser(userID int, request userPayload.UserUpdateRequest) (userPayload.UserResponse, error) {
	// Cek apakah user ada
	user, err := s.repository.FindByID(userID)

	if err != nil || user.ID == 0 {
		return userPayload.UserResponse{}, errors.New("user not found")
	}

	// Update user
	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Username != "" {
		user.Username = request.Username
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = request.Password
	}
	updateUser, err := s.repository.Update(user)
	if err != nil {
		return userPayload.UserResponse{}, err
	}

	return helpers.ToUserResponse(updateUser, ""), nil
}

func (s *newUserService) DeleteUser(userID int) (bool, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil || user.ID == 0 {
		return false, errors.New("user not found")
	}

	_, err = s.repository.Delete(user)
	if err != nil {
		return false, errors.New("telah terjadi kesalahan")
	}

	return true, nil
}

func (s *newUserService) GetUsers() ([]userPayload.UserResponse, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return helpers.ToUserResponses(users), err
	}

	return helpers.ToUserResponses(users), nil
}
