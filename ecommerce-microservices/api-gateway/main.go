package main

import (
	"bytes"
	"ecommerce-microservices/middleware"
	"io"
	"net/http"

	// Removed duplicate import
	// "api-gateway/middleware" // Removed duplicate import

	"github.com/gin-gonic/gin"
)

func proxyRequest(c *gin.Context, target string) {

	reqBody, _ := io.ReadAll(c.Request.Body)

	req, _ := http.NewRequest(
		c.Request.Method,
		target,
		bytes.NewBuffer(reqBody),
	)

	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func main() {

	r := gin.Default()

	// Public routes (Auth)
	r.POST("/register", func(c *gin.Context) {
		proxyRequest(c, "http://user-service:8001/register")
	})

	r.POST("/login", func(c *gin.Context) {
		proxyRequest(c, "http://user-service:8001/login")
	})

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/users", func(c *gin.Context) {
		proxyRequest(c, "http://user-service:8001/users")
	})

	protected.POST("/products", func(c *gin.Context) {
		proxyRequest(c, "http://product-service:8002/products")
	})

	protected.GET("/products", func(c *gin.Context) {
		proxyRequest(c, "http://product-service:8002/products")
	})

	protected.POST("/orders", func(c *gin.Context) {
		proxyRequest(c, "http://order-service:8003/orders")
	})

	r.Run(":8000")
}
