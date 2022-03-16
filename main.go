package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	time.Sleep(60 * time.Second)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, os.Getenv("PORTER_POD_IMAGE_TAG"))
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong2")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
