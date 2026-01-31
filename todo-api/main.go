package main

import (
	"todo-api/config"
	"todo-api/database"
	"todo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
