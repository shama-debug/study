package database

import (
	"fmt"
	"log"
	"os"

	"go-todo-app/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 他のパッケージからでもDBを操作できるようにグローバル関数化
var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("DB接続において.envファイルの読み込みに失敗しました")
	}
	// 環境変数から値を取得
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal("データベースの接続に失敗しました:", err)
	}

	DB = db
	log.Println("データベースの接続に成功しました (環境変数を使用)")

	// マイグレーション
	err = DB.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal("マイグレーションに失敗しました")
	}
	log.Println("マイグレーションに成功しました")
}
