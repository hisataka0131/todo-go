package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"todo-go/model"
)

var db *gorm.DB

func Init() {
	db, _ = gorm.Open("mysql", "root:@/go_sample?charset=utf8&parseTime=True&loc=Local")

	db.LogMode(true)

	db.AutoMigrate(&model.Todo{})

}

func Get() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
