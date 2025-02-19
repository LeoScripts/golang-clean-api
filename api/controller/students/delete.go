package students

import (
	"net/http"

	"golang-student-01/api/controller"
	"golang-student-01/entities/shared"
	student_usecase "golang-student-01/usecases/student"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	studentID, err := shared.GetUuidByStrings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("ERRO id invalido"))
		return
	}

	if err = student_usecase.Delete(studentID); err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError("Erro ao remover usuario, por favor tente mais tarde"))
		return
	}

	c.JSON(http.StatusOK, controller.NewResponseMessage("Estudante removido com sucesso"))
}
