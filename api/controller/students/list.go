package students

import (
	"net/http"

	student_usecase "golang-student-01/usecases/student"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	students, err := student_usecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, students)
}
