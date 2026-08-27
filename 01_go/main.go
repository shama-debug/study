package main

import (
	"go-todo-app/controllers"
	"go-todo-app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB接続とテーブル作成
	database.ConnectDB()
	// Ginルーターの初期化
	r := gin.Default()
	// ルーティングの設定(URLとコントローラを紐づける)
	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
	//サーバーの起動
	r.Run(":8080")
}
