package students

import (
	"golang-student-01/entities"
	"golang-student-01/entities/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var student entities.Student
	if err := c.Bind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro Payload vazio! por favor enviar os dados corretamente",
		})
		return
	}

	student.ID = shared.GetUuid() // criando com uuid //nao indicado pra escalar // usado para identificar micro serviços
	entities.StudentsMock = append(entities.StudentsMock, student)

	c.JSON(http.StatusCreated, student)
}
