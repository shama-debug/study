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
