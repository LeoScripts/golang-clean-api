package students

import (
	"golang-student-01/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, entities.StudentsMock)
}
