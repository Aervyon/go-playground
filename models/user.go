package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	IdBin    []byte `json:"idBin" gorm:"<-:create"`
	ID       string `json:"ID" gorm:"unique;primaryKey;<-:create"`
	Username string `json:"username"`
	Password string `json:"-"`
}
