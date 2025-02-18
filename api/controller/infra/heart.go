package heart

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HeartController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Deu bommmmmmmm",
	})

	c.Done()
}
