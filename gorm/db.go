package gorm

import (
	"github.com/jinzhu/gorm"
	"log"
)

const (
	Driver = "mysql"

	//DBUser = "root"
	//
	//DBPass = ""
	//
	//DBProtocol = "tcp(127.0.0.1:3306)"
	//
	//DBName = "go_sample"
)

func ConnectGorm() *gorm.DB {
	db, err := gorm.Open(Driver, "root:@/go_sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err.Error())
	}

	return db
}
