package utils

import (
	"fmt"

	"github.com/Aervyon/go-playground/models"

	"github.com/alexedwards/argon2id"
	"gorm.io/gorm"
)

func AuthUserPassword(db *gorm.DB, username, password string) (*models.UserModel, error) {
	user := &models.UserModel{}
	db.First(user, "username = ?", username)
	if user.Username == "" || user.Password == "" {
		return user, fmt.Errorf("user %v not found", username)
	}

	match, params, err := argon2id.CheckHash(password, user.Password)
	fmt.Println(params)
	if err != nil {
		return &models.UserModel{}, err
	}

	if !match {
		return &models.UserModel{}, fmt.Errorf("user failed authentication")
	}

	return user, nil
}
