# ex: arquitetura limpo(distribuindo responsabilidades)

com base no ex: 02

- movemos a nossa função controller para um arquivo `api/controller/infra/heart.go`

```go
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

```go
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

```go
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

```go
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

```go
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
```go
package students

type InputStudentDTO struct {
	FullName string `json:"full_name"`
	Age int `json:"age"`
}
```

os metodos que vao receber essa validação sao os que recebem dados externos como por exemplo o create e o update gerealmete


```go
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

```go
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

```go
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
```go
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

```go
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

```go
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

```go
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
```go
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

```go
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
```go
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
```go
package students

type InputStudentDto struct {
	FullName string `json:"full_name" validate:"required|min_len:3|max_len:150|string"`
	Age      int    `json:"age" validate:"required|int|min:3|max:80"`
}
```

- implementa no controller `api/controller/students/controller.go`

```go
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

###### neste outro exemplo vamos controlar a saida de dados

1. criamos nossa estrutura(struct) no DTO `api/controller/students/dto.go`
```go
package students

type InputStudentDto struct {
//............... codigo
}

// nossa nova estrutura(de saida)
// aqui controlamos  o que vai ser entregue ao usuario
// garantimos mais segurança e e passamos somente o que é preciso
type OutputStudentDto struct { //pra returno de um
	ID       uuid.UUID `json:"id"`
	FullName string `json:"full_name"`
	Age      int    `json:"age" `
}

type OutputStudentsDto struct { // pra retorno de varios
	Students []OutputStudentDto `json:"students"`
}

```

2. em `api/controller/students/validate.go`

```go

//condigo anterior(acima) ............

// ex: 01 do list
// func getOutputListStudents(students []entities.Student) (output OutputStudentsDto, err error) {
// 	for _, s := range students {
// 		output.Students = append(output.Students, OutputStudentDto{
// 			ID:       s.ID,
// 			FullName: s.FullName,
// 			Age:      s.Age,
// 		})
// 	}
// 	return output, err
// }

// ex: 02 do list
func getOutputListStudents(students []entities.Student) (output OutputStudentsDto, err error) {
	for _, s := range students {
		outputStudent, err := getOutputStudent(s)
		if err != nil {
			return output, err
		}
		output.Students = append(output.Students, outputStudent)
	}
	return output, err
}

func getOutputStudent(student entities.Student) (output OutputStudentDto, err error) {
	return OutputStudentDto{
		ID:       student.ID,
		FullName: student.FullName,
		Age:      student.Age,
	}, err
}

```
observe que no validate ja temos duas possibilidades de saida no caso de (1 item e de varios items, nesse caso de estudantes)

## Add Database e Repository

1. criar a estrutura do banco em `infra/database/database.go`
```go
package database

import (
	"golang-student-01/entities"
)

type Database struct {
	Conn              []*entities.Student
	StudentRepository entities.StudentRespository
}

func NewDatabase(conn []*entities.Student, sr entities.StudentRespository) *Database {
	return &Database{
		Conn:              conn,
		StudentRepository: sr,
	}
}
```

2. configurar o banco desejado em `infra/database/SEU_BANCO/SEU_BANCO.GO` (neste exemplo crie uma simulação de banco na memoria) ai e so trocar pelo que você desejar
```go
//banco em memoria somente pra testes

package memory

import (
	"golang-student-01/entities"

	"github.com/google/uuid"
)

// simulação de banco
var StudentsMemory = []*entities.Student{
	&entities.Student{uuid.Must(uuid.Parse("f69aa4db-7bdc-4d07-ad09-9f56ba2a7d6f")), "Alex", 38},
	&entities.Student{uuid.Must(uuid.Parse("98f6c71a-11f4-44b7-8712-32894fea0665")), "Bia", 32},
	&entities.Student{uuid.Must(uuid.Parse("3e35dc08-423b-4485-8783-8b374b1911f3")), "Carlos", 28},
}

// funçoes do banco
func FindAll() []*entities.Student {
	return StudentsMemory
}

func FindOne(id uuid.UUID) (student *entities.Student) {
	for _, sdt := range StudentsMemory {
		if sdt.ID == id {
			student = sdt
		}
	}
	return student
}

func Create(student *entities.Student) {
	StudentsMemory = append(StudentsMemory, student)
}

func Update(student *entities.Student) {
	var newStudents []*entities.Student

	for _, stud := range StudentsMemory {
		if student.ID == stud.ID {
			//atualiza
			newStudents = append(newStudents, student)
		} else {
			// mantem os existentes
			newStudents = append(newStudents, stud)
		}
	}
	StudentsMemory = newStudents
}

func Delete(id uuid.UUID) {
	var newStudents []*entities.Student
	for _, stdu := range StudentsMemory {
		if id != stdu.ID {
			newStudents = append(newStudents, stdu)
		}
	}
	StudentsMemory = newStudents
}

```
3. implementar a interface do repository dentro de `entities/SUA_ENTIDADE.GO` neste caso foi student.go
```go
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

type StudentRespository interface {
	List() []*Student  // so implementei o metodo list
}

```
4. criação do repository em `infra/database/repositories/SUA_ENTIDADE.GO` neste projeto acabei colocando o nome da pasta no singular `repository` mas deve ser colocado `repositories`
```go
package repository

import (
	"golang-student-01/entities"
	"golang-student-01/infra/database/memory"

	"github.com/google/uuid"
)

type StudentRepository struct {
	// Ctx context.Context
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		// Ctx: ctx,
	}
}

// ja inseri aqui um metodo
func (sr *StudentRepository) List() []*entities.Student {
	return memory.FindAll()
}
```
5. implementar a interface(o contrato) do usecase `entities/contract.go`
```go
package entities

type StudentUsecase interface {
	List() []*Student
}
```
6. fazer uso do nosso repository no usecase nexte exemplo em `usecases/student/list.go`
```go
package student

import (
	"golang-student-01/entities"
)

func (su *StudentUsecase) List() (students []*entities.Student) {
	return su.Database.StudentRepository.List()
}
```
7. instanciar nosso banco,usecase e controller  na incialização do serviço(server) em `api/api.go`
```go
package api

import (
	"golang-student-01/api/controller/students"
	"golang-student-01/infra/database"
	students_usecase "golang-student-01/usecases/student"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Engine            *gin.Engine
	DB                *database.Database // o database
	StudentController *students.StudentController
}

func NewService(db *database.Database) *Service {
	return &Service{
		Engine: gin.Default(),
		DB:     db, // o database
	}
}

func (s *Service) Start() {
	s.GetControllers() //depois aqui
	s.GetRoutes()
	s.Engine.Run(":7777")
}

//aqui
func (s *Service) GetControllers() {
	studentUsecase := students_usecase.NewStudentUsecase(s.DB)
	s.StudentController = students.NewStudentController(studentUsecase)
}
```
8. no arquivo root `main.go`
```go
package main

import (
	"golang-student-01/api"
	"golang-student-01/infra/database"
	"golang-student-01/infra/database/memory"
	"golang-student-01/infra/database/repository"
)

func main() {
	db := GetDatabase() // depois chamar aqui
	service := api.NewService(db)
	service.Start()
}

//aqui
func GetDatabase() *database.Database {
	db := memory.StudentsMemory
	studentRepository := repository.NewStudentRepository()
	return database.NewDatabase(db, studentRepository)
}
```
9. por fim usar o nosso usecase na controller neste caso em `api/controller/students/controller.go`
```go
package students

import (
	"net/http"

	"golang-student-01/api/controller"
	"golang-student-01/entities"
	"golang-student-01/usecases/student"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	StudentUsecase *student.StudentUsecase
}

func NewStudentController(su *student.StudentUsecase) *StudentController {
	return &StudentController{
		StudentUsecase: su,
	}
}

func (sc *StudentController) List(ctx *gin.Context) {
	students := sc.StudentUsecase.List()

	output, err := getOutputListStudents(students)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, output)
}
```
