package students

import (
	"golang-student-01/entities/shared"
	student_usecase "golang-student-01/usecases/student"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	studentID, err := shared.GetUuidByStrings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ERRO id invalido",
		})
		return
	}

	// regra de negocio
	if err = student_usecase.Delete(studentID); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao remover usuario, por favor tente mais tarde",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Estudante removido com sucesso",
	})

}
