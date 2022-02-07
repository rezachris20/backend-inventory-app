package service

import (
	"backend-inventory-app/entity"
	"backend-inventory-app/helpers"
	"backend-inventory-app/repository"
	"backend-inventory-app/web/role"
	"errors"
)

type RoleService interface {
	CreateNewRole(request role.CreateOrUpdateRoleRequest) (role.RoleResponses, error)
	UpdateRole(roleID int, request role.CreateOrUpdateRoleRequest) (role.RoleResponses, error)
	DeleteRole(roleID int) (bool, error)
	FindByID(roleID int) (role.RoleResponses, error)
	FindAll() ([]role.RoleResponses, error)
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return &roleService{roleRepository}
}

func (s *roleService) CreateNewRole(request role.CreateOrUpdateRoleRequest) (role.RoleResponses, error) {
	data := entity.Role{
		Name: request.Name,
	}

	role, err := s.roleRepository.Create(data)
	if err != nil {
		return helpers.ToRoleResponse(role), errors.New("create role failed")
	}

	return helpers.ToRoleResponse(role), nil
}

func (s *roleService) UpdateRole(roleID int, request role.CreateOrUpdateRoleRequest) (role.RoleResponses, error) {

	role, err := s.roleRepository.Role(roleID)
	if err != nil || role.ID == 0 {
		return helpers.ToRoleResponse(role), errors.New("role not found")
	}

	role.Name = request.Name

	update, err := s.roleRepository.Update(role)
	if err != nil {
		return helpers.ToRoleResponse(role), errors.New("failed to update role")
	}

	return helpers.ToRoleResponse(update), nil
}

func (s *roleService) DeleteRole(roleID int) (bool, error) {
	role, err := s.roleRepository.Role(roleID)
	if err != nil || role.ID == 0 {
		return false, errors.New("role not found")
	}

	_, err = s.roleRepository.Delete(role.ID)
	if err != nil {
		return false, errors.New("failed to delete role")
	}

	return true, nil
}

func (s *roleService) FindByID(roleID int) (role.RoleResponses, error) {
	role, err := s.roleRepository.Role(roleID)
	if err != nil || role.ID == 0 {
		return helpers.ToRoleResponse(role), errors.New("role not found")
	}

	return helpers.ToRoleResponse(role), nil
}

func (s *roleService) FindAll() ([]role.RoleResponses, error) {
	roles, err := s.roleRepository.Roles()
	if err != nil {
		return helpers.ToRoleResponses(roles), errors.New("record is empty")
	}

	return helpers.ToRoleResponses(roles), nil
}
