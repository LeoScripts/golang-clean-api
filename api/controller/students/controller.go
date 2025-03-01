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

func (sc *StudentController) Details(ctx *gin.Context) {
	var studentFound *entities.Student
	studentId, err := getInputId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("ID invalido"))
		return
	}

	studentFound, err = sc.StudentUsecase.SearchById(studentId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controller.NewResponseMessageError(err.Error()))
		return
	}

	output, err := getOutputStudent(studentFound)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}

	ctx.JSON(http.StatusOK, output)
}

func (sc *StudentController) Create(ctx *gin.Context) {
	input, err := getInputBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := sc.StudentUsecase.Create(input.FullName, input.Age)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, student)
}

func (sc *StudentController) Update(ctx *gin.Context) {

	studentID, err := getInputId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Erro id invalido"))
		return
	}
	input, err := getInputBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := sc.StudentUsecase.Update(studentID, input.FullName, input.Age)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	output, err := getOutputStudent(student)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}

	ctx.JSON(http.StatusOK, output)
}

// func (sc *StudentController) Delete(ctx *gin.Context) {

// 	studentID, err := getInputId(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("ERRO id invalido"))
// 		return
// 	}

// 	if err = student_usecase.Delete(studentID); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, controller.NewResponseMessageError("Erro ao remover usuario, por favor tente mais tarde"))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, controller.NewResponseMessage("Estudante removido com sucesso"))
// }
