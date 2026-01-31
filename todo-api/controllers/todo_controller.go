package controllers

import (
	"net/http"
	"strconv"
	"todo-api/database"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var todo models.Todo
	c.BindJSON(&todo)
	todo.UserID = userID

	database.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var todos []models.Todo
	var total int64

	database.DB.Where("user_id = ?", userID).Model(&models.Todo{}).Count(&total)
	database.DB.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&todos)

	c.JSON(http.StatusOK, gin.H{
		"data":  todos,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func UpdateTodo(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id := c.Param("id")

	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	if todo.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		return
	}

	c.BindJSON(&todo)
	database.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id := c.Param("id")

	var todo models.Todo
	database.DB.First(&todo, id)

	if todo.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		return
	}

	database.DB.Delete(&todo)
	c.Status(http.StatusNoContent)
}
