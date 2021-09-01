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

func GetUsers() (interface{}, error) {
	var users []models.Users

	if err := config.DB.Table("users").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}