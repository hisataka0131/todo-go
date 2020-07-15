package service

import (
	"github.com/jinzhu/gorm"
	m "todo-go/model"
)

// Todo作成
func insert(todo m.Todo, db *gorm.DB) {
	db.NewRecord(todo)
	db.Create(&todo)
}

// 全件取得
func findAll(db *gorm.DB) []m.Todo {
	var allTodo []m.Todo
	db.Find(&allTodo)
	return allTodo
}
