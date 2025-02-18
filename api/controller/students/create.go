package students

import (
	"golang-student-01/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var input InputStudentDto
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro Payload vazio! por favor enviar os dados corretamente",
		})
		return
	}

	student := entities.NewStudent(input.FullName, input.Age)
	entities.StudentsMock = append(entities.StudentsMock, *student)

	c.JSON(http.StatusCreated, student)
}
