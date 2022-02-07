package repository

import (
	"backend-inventory-app/entity"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role entity.Role) (entity.Role, error)
	Update(role entity.Role) (entity.Role, error)
	Delete(roleID int) (bool, error)
	Role(roleID int) (entity.Role, error)
	Roles() ([]entity.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) Create(role entity.Role) (entity.Role, error) {
	if result := r.db.Create(&role); result.Error != nil {
		return role, result.Error
	}

	return role, nil
}

func (r *roleRepository) Update(role entity.Role) (entity.Role, error) {
	if result := r.db.Updates(&role); result.Error != nil {
		return role, result.Error
	}

	return role, nil
}

func (r *roleRepository) Delete(roleID int) (bool, error) {
	var role entity.Role

	if result := r.db.Where("id = ? ", roleID).Delete(&role); result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (r *roleRepository) Role(roleID int) (entity.Role, error) {
	var role entity.Role

	if result := r.db.Where("id = ?", roleID).Find(&role); result.Error != nil {
		return role, result.Error
	}

	return role, nil
}

func (r *roleRepository) Roles() ([]entity.Role, error) {
	var roles []entity.Role

	if result := r.db.Find(&roles); result.Error != nil {
		return roles, result.Error
	}

	return roles, nil
}
