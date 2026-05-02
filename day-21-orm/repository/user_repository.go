package repository

import (
	"orm/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User

	result := r.DB.Find(&users)

	return users, result.Error
}

func (r *UserRepository) CreateUser(u *models.User) error {
	result := r.DB.Create(u)

	return result.Error
}

func (r *UserRepository) GetUserById(userId int) (models.User, error) {
	var user models.User

	result := r.DB.First(&user, "id = ?", userId)

	return user, result.Error
}

func (r *UserRepository) DeleteUser(userId int) error {
	result := r.DB.Delete(&models.User{}, userId)
	return result.Error
}

func (r *UserRepository) UpdateUser(userId int, user models.User) error {
	result := r.DB.Model(&models.User{}).Where("id = ?", userId).Updates(user)

	return result.Error
}
