package routes

import (
	"todo-api/controllers"
	"todo-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/todos", controllers.CreateTodo)
		auth.GET("/todos", controllers.GetTodos)
		auth.PUT("/todos/:id", controllers.UpdateTodo)
		auth.DELETE("/todos/:id", controllers.DeleteTodo)
	}
}
