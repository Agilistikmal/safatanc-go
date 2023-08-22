package database

import (
	"github.com/agilistikmal/safatanc-go/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func CreateConnection() *gorm.DB {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")+"?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Project{})
	if err != nil {
		panic(err)
	}
	DB = db

	return db
}
