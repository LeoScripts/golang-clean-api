# ex: arquitetura limpo(distribuindo responsabilidades)

com base no ex: 02 

- movemos a nossa função controller para um arquivo `api/controller/infra/heart.go`

```golang
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
```

- nossas rotas colocamos `api/route.go`

```golang
package api

import (
	heart "golang-student-01/api/controller/infra"
	"golang-student-01/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", heart.HeartController)
}
```

- a execução do servidor em `api/api.go`

```golang
package api

import "github.com/gin-gonic/gin"

type Service struct {
	*gin.Engine
}

func NewService() *Service {
	return &Service{
		gin.Default(),
	}
}

func (s *Service) Start() {
	s.GetRoutes()
	s.Engine.Run(":7777")
}
```

- e o `main.go` ficou assim somente chamndo nosso service

```golang
package main

import (
	"golang-student-01/api"
)

func main() {
	service := api.NewService()

	service.Start()
}
```

- passanos nossa entidade(regras de negocio) para `entities/student.go` e para esta pasta entities coloquei tambem a pasta `shared` ja que outras entidades poderam usar o id e tambem o nosso `mock` de estudantes dentro de student.go

```golang
package entities

import (
	"golang-student-01/entities/shared"

	"github.com/google/uuid"
)

type Student struct {
	ID       uuid.UUID `json:"_id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}

func NewStudent(fullName string, age int) *Student {
	return &Student{
		ID:       shared.GetUuid(),
		FullName: fullName,
		Age:      age,
	}
}

// mock
var StudentsMock = []Student{
	// substituimos o id pelo uuid
	Student{shared.GetUuid(), "Alex", 38},
	Student{shared.GetUuid(), "Bia", 32},
	Student{shared.GetUuid(), "Carlos", 28},
}
```
- em seguida fui removendo as funções controller para arquivos separando  dentor de `api/controller/students/NOME_DA_FUNÇÃO.go`, 
- atualizei as importações 
- e movi as rotas para o arquivo `api/route.go`

exemplo se a funcao de create coloquei em um arquivo create.go

o fluxo aqui foi o mesmo que fiz no em `api/controller/infra/heart.go`

ou seja para cada entidade crie uma pasta diferente