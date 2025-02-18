package students

import (
	"golang-student-01/entities/shared"
	student_usecase "golang-student-01/usecases/student"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var input InputStudentDto

	id := c.Params.ByName("id")
	studentID, err := shared.GetUuidByStrings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro id invalido",
		})
		return
	}

	if err = c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro Payload vazio! por favor enviar os dados corretamente",
		})
	}

	student, err := student_usecase.Update(studentID, input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, student)
}
