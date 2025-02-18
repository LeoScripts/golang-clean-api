package students

import (
	"golang-student-01/entities"
	"golang-student-01/entities/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var student entities.Student
	var newStudents []entities.Student

	id := c.Params.ByName("id")
	studentID, err := shared.GetUuidByStrings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ERRO id invalido",
		})
		return
	}

	// pegar o studante selecionado
	for _, stdu := range entities.StudentsMock {
		if studentID == stdu.ID {
			student = stdu
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "estudante não encontrado",
		})
		return
	}

	for _, stdu := range entities.StudentsMock {
		if stdu.ID != studentID {
			newStudents = append(newStudents, stdu)
		}
	}

	entities.StudentsMock = newStudents

	// retornar o item removomido
	c.JSON(http.StatusOK, gin.H{
		"message": "Estudante removido com sucesso",
		"student": student, // so pra testes
	})

}
