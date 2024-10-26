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

func (self *DB) Connect(host, user, password, dbname string, port int, sslmode, timeZone string) error {
    var err error
    connectionInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timeZone)
    self.DB, err = gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
    return err
}
