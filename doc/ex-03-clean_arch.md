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
- em seguida fui movendo as funções controller para arquivos separando  dentor de `api/controller/students/NOME_DA_FUNÇÃO.go`, 
- atualizei as importações 
- e movi as rotas para o arquivo `api/route.go`

exemplo se a funcao de create coloquei em um arquivo create.go

o fluxo aqui foi o mesmo que fiz para `api/controller/infra/heart.go` esse aquivo representa nossa rota de index ou home, por essa razao ficou na pasta infra dentro de controllers

ou seja para cada entidade crie uma pasta diferente


## colocando o DTO

estarei exibindo apenas 1 exemplo, ou seja faça nos outros que necessitarem tambem

a sigla DTO siguinifica "Data transfer object" = objeto de transferencia de dados
é aqui que dizemos quais dados viram e tambem validamos esses dados

- criei o DTO dentro de `api/controller/students/dto.go`
```golang
package students

type InputStudentDTO struct {
	FullName string `json:"full_name"`
	Age int `json:"age"`
}
```

os metodos que vao receber essa validação sao os que recebem dados externos como por exemplo o create e o update gerealmete


```golang
func Create(c *gin.Context) {
	var input InputStudentDto //estou aplicando aqui
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro Payload vazio! por favor enviar os dados corretamente",
		})
		return
	}

	student := entities.NewStudent(input.FullName, input.Age)
	entities.StudentsMock = append(entities.StudentsMock, *student)

	c.JSON(http.StatusCreated, student)
}
```

## adicioando Usecase(casos de uso)
nossa regra de negocio

⚠️⚠️⚠️ No clean arch nao podemos importar regras de outros arquivos que pertencem a outra implemtações mesmo que sejam iguais
o correto e duplicar o codigo mesmo ⚠️⚠️⚠️⚠️

estou mostrando somente um exemplo aqui continui a implemtação nos demais metodos

- criar dentro de `usecases/student/search.go`

obsever que dentro desse aquivo temos as regras trazida do controller `details.go`

```golang
package student

import (
	"errors"
	"golang-student-01/entities"
	"golang-student-01/entities/shared"

	"github.com/google/uuid"
)

func SearchById(id uuid.UUID) (student entities.Student, err error) {

	// regra: se o id recebido for igual a algum dentro do banco
	for _, stdu := range entities.StudentsMock {
		if id == stdu.ID {
			student = stdu
		}
	}

	// regra: se o o id so estudante encotrado for igual a 000000000000
	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Estudante não encontrado")
	}

	return student, err
}
```

- em `api/controller/students/details.go`

```golang
package students

import (
	"golang-student-01/entities"
	"golang-student-01/entities/shared"
	"net/http"

	student_usecase "golang-student-01/usecases/student"

	"github.com/gin-gonic/gin"
)

func Details(c *gin.Context) {
	var studentFound entities.Student
	id := c.Params.ByName("id")
	studentId, err := shared.GetUuidByStrings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID invalido",
		})
		return
	}

	// abistraimos para nossas regras para a camanda de usecases e ficou assim
	studentFound, err = student_usecase.SearchById(studentId)
	if err != nil {
		// vou abstrair isso em breve mas por enquanto vai ficar assim mesmo
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(), // a execução desse erro nos traz a msg que passamos dentro do usecase
		})
		return
	}

	c.JSON(http.StatusOK, studentFound)
}
```

## padronização do retorno

padronizando a mensagem de retorno e com isso desacoplando ainda mais o codigo
com a redução da dependencia do gin

lembrando que esses sao apenas alguns exemplos, implentei isso em varios locais dessa aplicação

- em `api/controller/response.go`
```golang
package controller

type Response struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Error string `json:"error"`
}

func NewResponseMessage(msg string) *Response {
	return &Response{
		Message: msg,
	}
}

func NewResponseMessageError(msg string) *ResponseError {
	return &ResponseError{
		Error: msg,
	}
}
```

em seguida fomos implemtando ex: dentro de `api/controller/students/create.go`

```golang
package students

import (
	"net/http"

	"golang-student-01/api/controller"
	student_usecase "golang-student-01/usecases/student"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var input InputStudentDto
	if err := c.Bind(&input); err != nil {

		// antes 
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message": "Erro Payload vazio! por favor enviar os dados corretamente",
		// })

		// agora
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Erro Payload vazio! por favor enviar os dados corretamente"))
		return
	}

	student, err := student_usecase.Create(input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, student)
}
```

outro exemplo de uso foi no heart `api/controller/infra/heart.go`

```golang
package heart

import (
	"net/http"

	"golang-student-01/api/controller"

	"github.com/gin-gonic/gin"
)

func HeartController(c *gin.Context) {
	// antes
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Deu bommmmmmmm",
	// })

	// agora
	c.JSON(http.StatusOK, controller.NewResponseMessage("Deu bommmmmmmm"))
}

```

## Refactor: movendo a logica dos controllers de student para um unico arquivo

nesta atualização coloquei as implemtaçoes dentro de metodos em um unico arquivo

criei um arquivo  `api/controller/students/controller.go` que é onde ficaram nossos metodos pra execução das acoes

```golang
package students

import (
	"golang-student-01/entities"
)

type StudentController struct {
	StudentUsecase *entities.Student
}

func NewStudentController(su *entities.Student) *StudentController {
	return &StudentController{
		StudentUsecase: su,
	}
}
```

depois movi todas as implemetações para dentro de metodos dentro do arquivo controller(nesse caso de students) como no exemplo abaixo que tinha que eu monstro o antes e o depois

- Antes <<<<<
```golang
package students

import (
	"net/http"

	student_usecase "golang-student-01/usecases/student"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {

	/// logica movida ----------------------------------
	students, err := student_usecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, students)
	/// -----------------------------------------------

}

```

pegue a logica ques esta marcada com comentario e coloquei no metodo List dentro de `api/controller/students/controller.go` dentro do metodo List

```golang
// .......codigo acima 

func (sc *StudentController) List(ctx *gin.Context) {
	students, err := student_usecase.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, students)
}
```

e fui fazendo isso nas demais funçoes e removendo os arquivos que agora nao eram mais necessarios

## validando os campos 

vou dar apenas alguns exemplos como base o restante e so olhar no codigo

estarei usando o pacote `github.com/gookit/validate`
baixe ai usando o `go get NOME_DO_PACOTE`

- criar o arquivo `api/controller/students/validate.go`
```golang
package students

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func getInputBody(ctx *gin.Context) (input InputStudentDto, err error) {
	// agora os dados passam primeiro por aqui 
	// depois eles seguem adiante na controller

	err = ctx.Bind(&input) // fazendo o bind
	if err != nil {
		return input, err
	}

	// validação de dados
	validation := validate.Struct(input)
	if !validation.Validate() {
		return input, validation.Errors
	}

	return input, err
}
```

- coloque as vaidaçoes no DTO como abaixo
```golang
package students

type InputStudentDto struct {
	FullName string `json:"full_name" validate:"required|min_len:3|max_len:150|string"`
	Age      int    `json:"age" validate:"required|int|min:3|max:80"`
}
```

- implementa no controller `api/controller/students/controller.go`

```golang
func (sc *StudentController) Create(ctx *gin.Context) {
	// Antes
	// var input InputStudentDto
	// if err := ctx.Bind(&input); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Erro Payload vazio! por favor enviar os dados corretamente"))
	// 	return
	// }

	// Depois -----------------------
	input, err := getInputBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}
	// ------------------------------

	// o restante continua do mesmo jeito
	student, err := student_usecase.Create(input.FullName, input.Age)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, student)
}
```



