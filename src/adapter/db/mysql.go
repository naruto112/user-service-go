package db

import (
	"fmt"
	"os"
	"user-service/src/adapter/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("MYSQLHOST")
	database = os.Getenv("MYSQL_DATABASE")
	user     = os.Getenv("MYSQLUSER")
	password = os.Getenv("MYSQLPASSWORD")
	port     = os.Getenv("MYSQLPORT")
)

// var (
// 	host     = "localhost"
// 	database = "user_service"
// 	user     = "root"
// 	password = "123"
// 	port     = "3307"
// )

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Mysqlconnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		checkError(err)
		panic("failed to connect database")
	} else {
		fmt.Printf("Connected to database %s\n", database)
		err = db.AutoMigrate(&entity.User{})
	}

	return db
}
