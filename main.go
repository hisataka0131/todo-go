package main

import (
	"todo-go/db"
	"todo-go/router"
)

func main() {

	// DB接続
	db.Init()
	defer db.Close()

	router.Router()
}
