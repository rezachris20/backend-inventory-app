package service

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/helpers"
	"backend-inventory-app/repository"
	"backend-inventory-app/web/user_role"
	"errors"
)

type UserRoleService interface {
	CreateNewUserRole(request user_role.CreateUserRoleRequest) (user_role.UserRoleResponse, error)
	UpdateUserRole(userRoleID int, request user_role.UpdateUserRoleRequest) (user_role.UserRoleResponse, error)
	DeleteUserRole(userRoleID int) (bool, error)
	GetUserRole(userRoleID int) (user_role.UserRoleResponse, error)
	GetUserRoles() ([]user_role.UserRoleResponse, error)
}

type userRoleService struct {
	userRoleRepository repository.UserRoleRepository
	userRepository     repository.UserRepository
	roleRepository     repository.RoleRepository
}

func NewUserRoleService(userRoleRepository repository.UserRoleRepository, userRepository repository.UserRepository, roleRepository repository.RoleRepository) UserRoleService {
	return &userRoleService{userRoleRepository, userRepository, roleRepository}
}

func (s *userRoleService) CreateNewUserRole(request user_role.CreateUserRoleRequest) (user_role.UserRoleResponse, error) {
	role, err := s.roleRepository.Role(request.RoleID)
	if err != nil || role.ID == 0 {
		return user_role.UserRoleResponse{}, errors.New("role not found")
	}

	user, err := s.userRepository.FindByID(request.UserID)
	if err != nil || user.ID == 0 {
		return user_role.UserRoleResponse{}, errors.New("user not found")
	}

	data := entity.UserRole{UserID: user.ID, RoleID: role.ID}

	userRole, err := s.userRoleRepository.Save(data)
	if err != nil {
		return user_role.UserRoleResponse{}, errors.New("something wrong")
	}

	response, _ := s.userRoleRepository.FindByID(userRole.ID)

	return helpers.ToUserRoleResponse(response), nil
}

func (s *userRoleService) UpdateUserRole(userRoleID int, request user_role.UpdateUserRoleRequest) (user_role.UserRoleResponse, error) {
	role, err := s.roleRepository.Role(request.RoleID)
	if err != nil || role.ID == 0 {
		return user_role.UserRoleResponse{}, errors.New("role not found")
	}

	user, err := s.userRepository.FindByID(request.UserID)
	if err != nil || user.ID == 0 {
		return user_role.UserRoleResponse{}, errors.New("user not found")
	}

	userRole, err := s.userRoleRepository.FindByID(userRoleID)
	if err != nil || userRole.ID == 0 {
		return user_role.UserRoleResponse{}, errors.New("user role not found")
	}

	data := entity.UserRole{ID: userRole.ID, UserID: user.ID, RoleID: role.ID}

	userRoleUpdated, err := s.userRoleRepository.Update(data)
	if err != nil {
		return user_role.UserRoleResponse{}, errors.New("something wrong")
	}

	response, _ := s.userRoleRepository.FindByID(userRoleUpdated.ID)

	return helpers.ToUserRoleResponse(response), nil
}

func (s *userRoleService) DeleteUserRole(userRoleID int) (bool, error) {
	role, err := s.userRoleRepository.FindByID(userRoleID)
	if err != nil || role.ID == 0 {
		return false, errors.New("user role not found")
	}

	_, err = s.userRoleRepository.Delete(role.ID)
	if err != nil {
		return false, errors.New("something wrong")
	}

	return true, nil
}

func (s *userRoleService) GetUserRole(userRoleID int) (user_role.UserRoleResponse, error) {
	userRole, err := s.userRoleRepository.FindByID(userRoleID)
	if err != nil || userRole.ID == 0 {
		return user_role.UserRoleResponse{}, errors.New("user role not found")
	}

	return helpers.ToUserRoleResponse(userRole), nil
}

func (s *userRoleService) GetUserRoles() ([]user_role.UserRoleResponse, error) {
	userRoles, err := s.userRoleRepository.FindByAll()
	if err != nil {
		return helpers.ToUserRoleResponses(userRoles), nil
	}

	return helpers.ToUserRoleResponses(userRoles), nil
}
