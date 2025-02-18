package students

import (
	"golang-student-01/entities"
	"golang-student-01/entities/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Details(c *gin.Context) {
	var student entities.Student
	id := c.Params.ByName("id")
	studentId, err := shared.GetUuidByStrings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID invalido",
		})
		return
	}

	for _, stdu := range entities.StudentsMock {
		if studentId == stdu.ID {
			student = stdu
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Estudante não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}
