package students

import (
	"net/http"

	"golang-student-01/api/controller"
	"golang-student-01/entities"
	student_usecase "golang-student-01/usecases/student"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	StudentUsecase *entities.StudentUsecaseContract
}

func NewStudentController(su *entities.StudentUsecaseContract) *StudentController {
	return &StudentController{
		StudentUsecase: su,
	}
}

func (sc *StudentController) List(ctx *gin.Context) {
	students, err := student_usecase.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, students)
}

func (sc *StudentController) Create(ctx *gin.Context) {
	var input InputStudentDto
	if err := ctx.Bind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Erro Payload vazio! por favor enviar os dados corretamente"))
		return
	}

	student, err := student_usecase.Create(input.FullName, input.Age)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, student)
}
