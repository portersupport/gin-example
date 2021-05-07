package main

import (
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  r := gin.Default()

  r.GET("/ping", func(c *gin.Context) {
    fmt.Println("CLIENT IP IS", c.ClientIP())

    c.String(http.StatusOK, "pong")
  })

  return r
}

func main() {
  r := setupRouter()
  // Listen and Server in 0.0.0.0:8080
  r.Run(":8080")
}