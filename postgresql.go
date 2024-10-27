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

/* For registration/login accept this fields
*  When user registered or login - return token 
    For registration:
        Create token
        Save token in cache ( userId - token)
        Return token
    For login:
        Get or create token from cache 
        Save in cache if token was created
        Return token
type UserRegistration struct {
    Name string          `json:"name" gorm:"type:text"`
    Email string         `json:"email" gorm:"type:text unique"`
    Password string      `json:"password" gorm:"type:text"`
}
*/

type DB struct {
    DB *gorm.DB
}

func DBNew() *DB {
    obj := new(DB)
    return obj
}

func (self *DB) Connect(host, user, password, dbname string, port int, sslmode, timeZone string) error {
    var err error
    connectionInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timeZone)
    self.DB, err = gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
    return err
}

func (self *DB) CreateNewUser(user UserDB) (uint, error) {
    // _, err := self.GetUserByEmail(user.Email)

    if self.isUserEmailExist(user.Email) {
        return 0, fmt.Errorf("User with email: '%s' exist\n", user.Email)
    }

    err := self.DB.Model(&UserDB{}).Create(&user).Error
    return user.Id, err
}

func (self *DB) GetUserById(id uint) (UserDB, error) {
    var user UserDB
    err := self.DB.First(&user, id).Error
    return user, err
}

func (self *DB) GetUserByEmail(email string) (UserDB, error) {
    var user UserDB
    err := self.DB.Model(&UserDB{}).First(&user, "email = ?", email).Error
    return user, err
}

func (self *DB) isUserEmailExist(email string) bool {
    _, err := self.GetUserByEmail(email)

    if err == gorm.ErrRecordNotFound {
        return false
    }

    return true
}
