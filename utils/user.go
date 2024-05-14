package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type User struct {
	IdBin    []byte `json:"idBin"`
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type UserModel struct {
	gorm.Model
	IdBin    []byte `json:"idBin" gorm:"<-:create"`
	ID       string `json:"ID" gorm:"unique;primaryKey;<-:create"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func (*UserModel) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateUser(db *gorm.DB, username, password string) (*UserModel, error) {
	id := ulid.Make()

	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)

	if err != nil {
		return &UserModel{}, err
	}

	user := &UserModel{IdBin: id.Bytes(), ID: id.String(), Password: hash, Username: username}
	transaction := db.Model(&UserModel{}).Create(user)
	if transaction.Error != nil {
		return &UserModel{}, transaction.Error
	}
	log.Println("Created user", username)
	return user, nil
}

func AuthUserPassword(db *gorm.DB, username, password string) (*UserModel, error) {
	user := &UserModel{}
	db.First(user, "username = ?", username)
	if user.Username == "" || user.Password == "" {
		return user, fmt.Errorf("user %v not found", username)
	}

	match, params, err := argon2id.CheckHash(password, user.Password)
	fmt.Println(params)
	if err != nil {
		return &UserModel{}, err
	}

	if !match {
		return &UserModel{}, fmt.Errorf("user failed authentication")
	}

	return user, nil
}
