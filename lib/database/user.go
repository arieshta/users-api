package database

import (
	"users-api/models"
)

func CreateUser(user *models.Users) error {
	if err := config.DB.Table("users").Create(&user).Error; err != nil {
		return err
	}
	return nil
}