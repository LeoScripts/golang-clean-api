package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID       int    `json:"_id"`
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
}

// mock
var Students = []Student{
	Student{1, "Alex", 38},
	Student{2, "Bia", 32},
	Student{3, "Carlos", 28},
}

func heartController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Deu bommmmmmmm",
	})

	c.Done()
}

func routeGetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func main() {
	service := gin.Default()

	getRoutes(service)

	service.Run(":7777")

}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", heartController)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", routeGetStudents)
	return c
}
