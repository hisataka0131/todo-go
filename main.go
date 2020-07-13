package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	db "todo-go/gorm"
	"todo-go/model"
)

func main() {

	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "こんにちは",
		})
	})

	r.Run()

	db := db.ConnectGorm()
	defer db.Close()
	db.AutoMigrate(&model.Todo{})
}
