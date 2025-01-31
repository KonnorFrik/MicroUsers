package main

import (
	"golang.org/x/crypto/bcrypt"
)

type UserIn struct {
    Email string         `json:"email" gorm:"type:text unique"`
    Password string      `json:"password" gorm:"type:text"` // On register - no hashed
}

type UserOut struct {
    Email string         `json:"email" gorm:"type:text unique"`
}

type UserDB struct {
    Email string         `json:"email" gorm:"type:text unique"`
    Password string      `json:"password" gorm:"type:text"` // On register - no hashed
}

/*
Encrypt a password with bcrypt and with cost - 8
*/
func CryptPassword(password string) (string, error) {
    hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
    return string(hashPassword), err
}

