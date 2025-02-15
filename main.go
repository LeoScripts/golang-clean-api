package main

import (
	"net/http"
	"strconv"

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

func getStudentsController(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func createStudentController(c *gin.Context) {
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

func updateStudentController(c *gin.Context) {
	var student Student
	var studentTemp Student
	var newStudents []Student
	// params := c.Params //pega todos os paramentros

	// id := c.Params.ByName("id") //pega um paramentro id
	// studentID, err := strconv.Atoi(id)

	studentID, err := strconv.Atoi(c.Params.ByName("id")) //melhora a performace (pois reduzo a qtd de variaveis)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro id invalido",
		})
		return
	}

	if err = c.Bind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro Payload vazio! por favor enviar os dados corretamente",
		})
	}

	for _, sdt := range Students {
		if sdt.ID == studentID {
			studentTemp = sdt
		}
	}

	if studentTemp.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id nao encotrado",
		})
		return
	}

	studentTemp.FullName = student.FullName
	studentTemp.Age = student.Age

	for _, stud := range Students {
		if studentID == stud.ID {
			newStudents = append(newStudents, studentTemp)
		} else {
			newStudents = append(newStudents, stud)
		}
	}

	Students = newStudents

	c.JSON(http.StatusOK, studentTemp)
}

func main() {
	service := gin.Default()

	getRoutes(service)

	service.Run(":7777")

}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", heartController)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", getStudentsController)
	groupStudents.POST("/", createStudentController)
	groupStudents.PUT("/:id", updateStudentController)

	return c
}
