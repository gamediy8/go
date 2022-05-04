package data

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UserDb *gorm.DB

func init() {
	dsn := "page:admin888...@tcp(127.0.0.1:3336)/lottery_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("user db error")
	}
	UserDb = db
	fmt.Println("user db init success")

}
