package repository

import (
	"backend-inventory-app/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRoleRepository interface {
	Save(userRole entity.UserRole) (entity.UserRole, error)
	Update(userRole entity.UserRole) (entity.UserRole, error)
	Delete(userRoleID int) (bool, error)
	FindByID(userRoleID int) (entity.UserRole, error)
	FindByAll() ([]entity.UserRole, error)
}

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{db}
}

func (r *userRoleRepository) Save(userRole entity.UserRole) (entity.UserRole, error) {

	if err := r.db.Create(&userRole).Error; err != nil {
		return userRole, err
	}

	return userRole, nil
}

func (r *userRoleRepository) Update(userRole entity.UserRole) (entity.UserRole, error) {
	if err := r.db.Debug().Updates(&userRole).Error; err != nil {
		return userRole, err
	}

	return userRole, nil
}

func (r *userRoleRepository) Delete(userRoleID int) (bool, error) {
	var userRole entity.UserRole

	if err := r.db.Where("id = ?", userRoleID).Delete(&userRole).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *userRoleRepository) FindByID(userRoleID int) (entity.UserRole, error) {
	var userRole entity.UserRole

	if err := r.db.Preload(clause.Associations).Where("id = ?", userRoleID).Find(&userRole).Error; err != nil {
		return userRole, err
	}

	return userRole, nil
}

func (r *userRoleRepository) FindByAll() ([]entity.UserRole, error) {
	var userRoles []entity.UserRole

	if err := r.db.Debug().Preload(clause.Associations).Find(&userRoles).Error; err != nil {
		return userRoles, err
	}

	return userRoles, nil
}
