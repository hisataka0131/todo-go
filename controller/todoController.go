package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"todo-go/model"
)

type TodoHandler struct {
	Db *gorm.DB
}

// 一覧表示
func (handler *TodoHandler) GetAll(c *gin.Context) {
	var todos []model.Todo
	handler.Db.Find(&todos)
	c.HTML(http.StatusOK, "index.html", gin.H{"todos": todos})
}

// 新規作成
func (handler *TodoHandler) Create(c *gin.Context) {
	text, _ := c.GetPostForm("text")
	handler.Db.Create(&model.Todo{Text: text, IsDone: false})
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 編集画面
func (handler *TodoHandler) Edit(c *gin.Context) {
	todo := model.Todo{}
	id := c.Param("id")
	handler.Db.First(&todo, id)
	c.HTML(http.StatusOK, "edit.html", gin.H{"todo": todo})
}

// 更新
func (handler *TodoHandler) Update(c *gin.Context) {
	todo := model.Todo{}
	id := c.Param("id")
	text, _ := c.GetPostForm("text")
	handler.Db.First(&todo, id)
	todo.Text = text
	handler.Db.Save(&todo)
	c.Redirect(http.StatusMovedPermanently, "/")
}

// タスク更新
func (handler *TodoHandler) IsDone(c *gin.Context) {
	todo := model.Todo{}
	id := c.Param("id")
	handler.Db.First(&todo, id)
	if todo.IsDone {
		todo.IsDone = false
	} else {
		todo.IsDone = true
	}
	handler.Db.Save(&todo)
	c.Redirect(http.StatusMovedPermanently, "/"+id)
}

// 削除
func (handler *TodoHandler) Delete(c *gin.Context) {
	todo := model.Todo{}
	id := c.Param("id")
	handler.Db.First(&todo, id)
	handler.Db.Delete(&todo)
	c.Redirect(http.StatusMovedPermanently, "/")
}
