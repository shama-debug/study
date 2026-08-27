package controllers

import (
	"go-todo-app/database"
	"go-todo-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODOリスト一覧
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// 新規TODO作成
func CreateTodo(c *gin.Context) {
	var input struct {
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := models.Todo{Title: input.Title}
	database.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// TODO編集
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "指定されたTodoが見つかりません。"})
		return
	}
	// 入力の型を定義
	var input struct {
		Title  string `json:"title"`
		Status string `json:"status"`
		Memo   string `json:"memo"`
	}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// DBでidに紐づくTODOを更新
	database.DB.Model(&todo).Updates(todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// TODO削除
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "指定されたTodoが見つかりません。"})
		return
	}
	// DBからTodoを削除
	database.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"message": "Todoを削除しました"})
}
