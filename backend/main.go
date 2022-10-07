package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, "hello home")
}

func main() {
	router := gin.Default()
	router.GET("/", home)
	router.GET("/api", hello)

	router.Run()
}
