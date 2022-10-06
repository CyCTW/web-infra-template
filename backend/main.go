package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}
func main() {
	router := gin.Default()
	router.GET("/api", hello)

	router.Run("localhost:8081")
}
