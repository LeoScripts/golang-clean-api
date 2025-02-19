package students

import (
	"net/http"

	"golang-student-01/api/controller"
	student_usecase "golang-student-01/usecases/student"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var input InputStudentDto
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Erro Payload vazio! por favor enviar os dados corretamente"))
		return
	}

	student, err := student_usecase.Create(input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, student)
}
