package database

import (
	"users-api/models"
	"users-api/config"
)

func CreateUser(user *models.Users) error {
	if err := config.DB.Table("users").Create(&user).Error; err != nil {
		return err
	}
	return nil
}