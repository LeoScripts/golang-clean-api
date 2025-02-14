package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func heartController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Deu bommmmmmmm",
	})

	c.Done()
}

func main() {
	service := gin.Default()

	getRoutes(service)

	service.Run(":7777")

}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", heartController)
	return c
}
