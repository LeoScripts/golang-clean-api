package heart

import (
	"net/http"

	"golang-student-01/api/controller"

	"github.com/gin-gonic/gin"
)

func HeartController(c *gin.Context) {
	c.JSON(http.StatusOK, controller.NewResponseMessage("Deu bommmmmmmm"))
}
