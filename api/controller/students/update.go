package students

import (
	"golang-student-01/entities"
	"golang-student-01/entities/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var input InputStudentDto
	var studentTemp entities.Student
	var newStudents []entities.Student

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

	for _, sdt := range entities.StudentsMock {
		if sdt.ID == studentID {
			studentTemp = sdt
		}
	}

	if studentTemp.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id nao encotrado",
		})
		return
	}

	studentTemp.FullName = input.FullName
	studentTemp.Age = input.Age

	for _, stud := range entities.StudentsMock {
		if studentID == stud.ID {
			newStudents = append(newStudents, studentTemp)
		} else {
			newStudents = append(newStudents, stud)
		}
	}

	entities.StudentsMock = newStudents
	c.JSON(http.StatusOK, studentTemp)
}
