package router

import (
	"github.com/gin-gonic/gin"
	"todo-go/controller"
	"todo-go/db"
)

func Router() {

	router := gin.Default()

	// htmlファイルを関係づける
	router.LoadHTMLGlob("templates/*.html")

	handler := controller.TodoHandler{
		Db: db.Get(),
	}

	// 各パスにGET/POSTメソッドでリクエストされた時のレスポンスを登録
	router.GET("/", handler.GetAll)
	router.POST("/", handler.Create)
	router.GET("/:id", handler.Edit)
	router.POST("/update/:id", handler.Update)
	router.POST("/delete/:id", handler.Delete)
	router.POST("/done/:id", handler.IsDone)

	router.Run()
}
