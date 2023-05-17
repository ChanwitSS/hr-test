package services

import (
	"errors"
	"hr/internal/models"
	"hr/pkg/app/util"
)

func Register(user models.User) (*models.User, error) {
	result, err := models.CreateUser(models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,

		CreatedAt: util.GetLocalTime("Bangkok"),
		UpdatedAt: util.GetLocalTime("Bangkok"),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Login(username string, password string, email string) (bool, error) {
	result, err := models.CheckAuthUser(username, password, email)
	if err != nil {
		return false, err
	}

	if result.Password != password {
		err := errors.New("Password not match.\n")
		return false, err
	}

	return true, nil
}
