package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

const (
	user = "dev"
	pass = "admin"
	host = "mysql"
	port = "3306"
)

func Connect() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/imm?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
