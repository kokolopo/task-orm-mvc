package config

import (
	"task-orm-mvc/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {

	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/dbName?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	user := models.User{}
	tiket := models.Tiket{}
	transaksi := models.Transaksi{}

	DB.AutoMigrate(&user)
	DB.AutoMigrate(&tiket)
	DB.AutoMigrate(&transaksi)
}
