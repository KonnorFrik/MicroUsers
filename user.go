package main

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

const (
    tokenLength = 6
)

type UserDB struct {
    Id uint              `json:"id" gorm:"primaryKey"`
    Name string          `json:"name" gorm:"type:text"`
    Password string      `json:"password" gorm:"type:text"` // On register - no hashed
    Email string         `json:"email" gorm:"type:text unique"`
}

func CryptPassword(user UserDB) (string, error) {
    hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
    return string(hashPassword), err
}

func GenerateToken() (string, error) {
    byteBuffer := make([]byte, tokenLength)
    _, err := rand.Read(byteBuffer)

    if err != nil {
        return "", err
    }

    result := base64.StdEncoding.EncodeToString(byteBuffer)
    return result, nil
}
