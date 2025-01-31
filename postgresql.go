package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
    DB_DEFAULT_HOST = "postgres"
    DB_DEFAULT_USER = "postgres"
    DB_DEFAULT_PASSWORD = "admin"
    DB_DEFAULT_DBNAME = "microusers"
    DB_DEFAULT_PORT = 5432
    DB_DEFAULT_SSLMODE = "disable"
    DB_DEFAULT_TIMEZONE = "Asia/Yekaterinburg"
)

type DB struct {
    DB *gorm.DB
}

func DBNew() *DB {
    obj := new(DB)
    return obj
}

/*
Connect to postgres DB with given settings
*/
func (self *DB) Connect(host, user, password, dbname string, port int, sslmode, timeZone string) error {
    // commented info uses a different from default database
    // TODO: find a way for safe create database if not exist and reconnect using new db
    // connectionInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timeZone)

    var err error
    connectionInfo := fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=%s TimeZone=%s", host, user, password, port, sslmode, timeZone)
    self.DB, err = gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})

    return err
}

/*
Create a new user DB
Password in user struct must be already encrypted
*/
func (self *DB) CreateNewUser(user UserDB) error {
    if self.IsUserEmailExist(user.Email) {
        return fmt.Errorf("User with email: '%s' already exist\n", user.Email)
    }

    err := self.DB.Model(&UserDB{}).Create(&user).Error
    return err
}

// func (self *DB) GetUserById(id uint) (UserDB, error) {
//     var user UserDB
//     err := self.DB.First(&user, id).Error
//     return user, err
// }

/*
Get a user from DB by email
*/
func (self *DB) GetUserByEmail(email string) (UserDB, error) {
    var user UserDB
    err := self.DB.Model(&UserDB{}).First(&user, "email = ?", email).Error
    return user, err
}

/*
Check if email exist in DB
*/
func (self *DB) IsUserEmailExist(email string) bool {
    _, err := self.GetUserByEmail(email)

    if err == gorm.ErrRecordNotFound {
        return false
    }

    return true
}
