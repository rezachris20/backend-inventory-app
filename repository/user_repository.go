package repository

import (
	"backend-inventory-app/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(user entity.User) (bool, error)
	FindByID(userID int) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Delete(user entity.User) (bool, error) {
	err := r.db.Where("id = ? ", user.ID).Delete(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *userRepository) FindByID(userID int) (entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}
