package students

import (
	"net/http"

	"golang-student-01/api/controller"
	"golang-student-01/entities"
	"golang-student-01/entities/shared"
	student_usecase "golang-student-01/usecases/student"

	"github.com/gin-gonic/gin"
)

func Details(c *gin.Context) {
	var studentFound entities.Student
	id := c.Params.ByName("id")
	studentId, err := shared.GetUuidByStrings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("ID invalido"))
		return
	}

	studentFound, err = student_usecase.SearchById(studentId)
	if err != nil {
		c.JSON(http.StatusNotFound, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, studentFound)
}
