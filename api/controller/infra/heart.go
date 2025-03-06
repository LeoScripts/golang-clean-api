package heart

import (
	"net/http"

	"github.com/LeoScripts/golang-clean-api/api/controller"

	"github.com/gin-gonic/gin"
)

func HeartController(c *gin.Context) {
	c.JSON(http.StatusOK, controller.NewResponseMessage("Deu bommmmmmmm"))
}
