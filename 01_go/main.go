package main

import (
	"go-todo-app/controllers"
	"go-todo-app/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB接続と .env の読み込み
	database.ConnectDB()
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET 環境変数を設定してください")
	}
	// Ginルーターの初期化
	r := gin.Default()
	// ルーティングの設定(URLとコントローラを紐づける)
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)

	todos := r.Group("/todos")
	todos.Use(controllers.AuthRequired())
	todos.GET("", controllers.GetTodos)
	todos.POST("", controllers.CreateTodo)
	todos.PUT("/:id", controllers.UpdateTodo)
	todos.DELETE("/:id", controllers.DeleteTodo)
	//サーバーの起動
	r.Run(":8080")
}
