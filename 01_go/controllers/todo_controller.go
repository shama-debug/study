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
	database.DB.Where("user_id = ?", currentUserID(c)).Find(&todos)
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
	todo := models.Todo{Title: input.Title, UserID: currentUserID(c)}
	database.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// TODO編集
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.Where("id = ? AND user_id = ?", id, currentUserID(c)).First(&todo).Error; err != nil {
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
	if err := database.DB.Model(&todo).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Todoの更新に失敗しました"})
		return
	}
	if err := database.DB.First(&todo, todo.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Todoの取得に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// TODO削除
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.Where("id = ? AND user_id = ?", id, currentUserID(c)).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "指定されたTodoが見つかりません。"})
		return
	}
	// DBからTodoを削除
	database.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"message": "Todoを削除しました"})
}
