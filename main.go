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

func GetStudentsController(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func CreateStudentController(c *gin.Context) {
	var student Student
	if err := c.Bind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro Payload vazio! por favor enviar os dados corretamente",
		})
		return
	}
	// student.ID = len(Students) + 1 //funciona mais nao e bom
	student.ID = Students[len(Students)-1].ID + 1 // a melhor opção seria deixar pro banco fazer isso
	Students = append(Students, student)

	c.JSON(http.StatusCreated, student)
}

func main() {
	service := gin.Default()

	getRoutes(service)

	service.Run(":7777")

}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", heartController)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", GetStudentsController)
	groupStudents.POST("/", CreateStudentController)

	return c
}
