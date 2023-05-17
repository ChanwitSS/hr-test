package models

import (
	"time"
)

type User struct {
	UserId   int    `gorm:"primary_key" json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	// ProfileImg string `json:"profile_img"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUser(user User) (*User, error) {
	if err := db.Table("users").Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUser(userId string) (*User, error) {
	var user User
	err := db.
		Table("users").
		Where("users.user_id = ?", userId).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CheckAuthUser(username string, password string, email string) (*User, error) {
	var user User
	err := db.
		Table("users").
		Where(&User{
			Username: username,
		}).Or(&User{
		Password: password,
	}).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
