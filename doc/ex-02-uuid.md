# exemplo 02 usando o uuid

- shared/id.go
```golang
package shared

import (
	"log"

	"github.com/google/uuid"
)

// cria um uuid
func GetUuid() uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Fatal error", err)
	}
	return uuid
}

// transforma a string que recebemos para uuid
func GetUuidByStrings(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

// usando pra passar valores ids nulos ou zeros
func GetUuidEmpty() uuid.UUID {
	return uuid.Nil
}

```

- main.go
```golang
package main

import (
	"golang-student-01/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Student struct {
	ID       uuid.UUID `json:"_id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}

// mock
var Students = []Student{
	// substituimos o id pelo uuid
	Student{shared.GetUuid(), "Alex", 38},
	Student{shared.GetUuid(), "Bia", 32},
	Student{shared.GetUuid(), "Carlos", 28},
}

func heartController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Deu bommmmmmmm",
	})

	c.Done()
}

func getStudentsController(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func createStudentController(c *gin.Context) {
	var student Student
	if err := c.Bind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro Payload vazio! por favor enviar os dados corretamente",
		})
		return
	}

	student.ID = shared.GetUuid() // criando com uuid //nao indicado pra escalar // usado para identificar micro serviços
	Students = append(Students, student)

	c.JSON(http.StatusCreated, student)
}

func updateStudentController(c *gin.Context) {
	var student Student
	var studentTemp Student
	var newStudents []Student

	// studentID, err := strconv.Atoi(c.Params.ByName("id")) //melhora a performace (pois reduzo a qtd de variaveis)
	id := c.Params.ByName("id")
	studentID, err := shared.GetUuidByStrings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro id invalido",
		})
		return
	}

	if err = c.Bind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro Payload vazio! por favor enviar os dados corretamente",
		})
	}

	for _, sdt := range Students {
		if sdt.ID == studentID {
			studentTemp = sdt
		}
	}

	// if studentTemp.ID == 0 {
	if studentTemp.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id nao encotrado",
		})
		return
	}

	studentTemp.FullName = student.FullName
	studentTemp.Age = student.Age

	for _, stud := range Students {
		if studentID == stud.ID {
			newStudents = append(newStudents, studentTemp)
		} else {
			newStudents = append(newStudents, stud)
		}
	}

	Students = newStudents
	// geralmento so se retorna uma msg dizendo que foi atualizado
	c.JSON(http.StatusOK, studentTemp) // estamos mostrando o studante , mas isso nao e necessario
}

func deleteStudentController(c *gin.Context) {
	var student Student
	var newStudents []Student

	id := c.Params.ByName("id")
	studentID, err := shared.GetUuidByStrings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ERRO id invalido",
		})
		return
	}

	// pegar o studante selecionado
	for _, stdu := range Students {
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

	for _, stdu := range Students {
		if stdu.ID != studentID {
			newStudents = append(newStudents, stdu)
		}
	}

	Students = newStudents

	// retornar o item removomido
	c.JSON(http.StatusOK, gin.H{
		"message": "Estudante removido com sucesso",
		"student": student, // so pra testes
	})

}

func getStudentsByIdController(c *gin.Context) {
	var student Student
	id := c.Params.ByName("id")
	studentId, err := shared.GetUuidByStrings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID invalido",
		})
		return
	}

	for _, stdu := range Students {
		if studentId == stdu.ID {
			student = stdu
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Estudante não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}
func main() {
	service := gin.Default()

	getRoutes(service)

	service.Run(":7777")

}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", heartController)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", getStudentsController)
	groupStudents.POST("/", createStudentController)
	groupStudents.PUT("/:id", updateStudentController)
	groupStudents.DELETE("/:id", deleteStudentController)
	groupStudents.GET("/:id", getStudentsByIdController)

	return c
}

```