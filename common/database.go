package common

import (
	"Vnollx/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// func InitDB() *gorm.DB {
func InitDB() {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "test_db"
	username := "root"
	password := "ab147890"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, _ := gorm.Open(driverName, args)
	/*if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}*/

	//迁移
	db.AutoMigrate(&model.User{})

	//初始化后的db赋给全局变量DB，没赋值的化DB就是一个nil
	DB = db
	//return db
}
func GetDB() *gorm.DB {
	return DB
}
