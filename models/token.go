package models

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Token struct {
	ID    string `gorm:"primarykey;unique"`
	UID   string `gorm:"primarykey"`
	Hash  string `json:"-"`
	Token string `gorm:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

/*
* returns a 4 byte timestamp
 */
func MakeTimestamp(t time.Time) []byte {
	unix := t.Unix()
	bytes := make([]byte, 4)
	bytes[0] = byte(unix >> 24)
	bytes[1] = byte(unix >> 16)
	bytes[2] = byte(unix >> 8)
	bytes[3] = byte(unix)
	return bytes
}

/*
	A simple function to make just the token itself

returns a 20 byte token in string and []byte form
*/
func MakeToken(t time.Time) (string, []byte) {
	timestamp := MakeTimestamp(t)
	random := make([]byte, 16)
	_, err := rand.Read(random)
	if err != nil {
		panic(err)
	}

	token := timestamp
	token = append(token, random...)

	return hex.EncodeToString(token), token
}

func NewToken(uid string) *Token {
	// make id
	id := ulid.Make()

	// Make the token. 4 bytes of time and 20 of entropy
	token, _ := MakeToken(time.Now())

	// hash the token
	hash, err := argon2id.CreateHash(token, argon2id.DefaultParams)
	if err != nil {
		panic(err)
	}

	t := &Token{
		Token:     token,
		ID:        id.String(),
		UID:       uid,
		Hash:      hash,
		CreatedAt: time.Now(),
	}

	t.UpdatedAt = t.CreatedAt
	return t
}
