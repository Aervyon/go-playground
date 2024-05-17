package database

// Package database is a helper package to do things like creating models, and deleting data (if complex)

// If operations are not complex, they should not be added to the database pacakge.

import (
	"log"

	"github.com/Aervyon/go-playground/models"
	"github.com/alexedwards/argon2id"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, username, password string) (*models.UserModel, error) {
	id := ulid.Make()

	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)

	if err != nil {
		return &models.UserModel{}, err
	}

	user := &models.UserModel{IdBin: id.Bytes(), ID: id.String(), Password: hash, Username: username}
	transaction := db.Model(&models.UserModel{}).Create(user)
	if transaction.Error != nil {
		return &models.UserModel{}, transaction.Error
	}
	log.Println("Created user", username)
	return user, nil
}
