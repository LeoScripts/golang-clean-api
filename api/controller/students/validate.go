package students

import (
	"golang-student-01/entities"
	"golang-student-01/entities/shared"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

func getInputBody(ctx *gin.Context) (input InputStudentDto, err error) {
	err = ctx.Bind(&input)
	if err != nil {
		return input, err
	}

	validation := validate.Struct(input)
	if !validation.Validate() {
		return input, validation.Errors
	}

	return input, err
}

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

func getInputId(ctx *gin.Context) (id uuid.UUID, err error) {
	inputId := ctx.Params.ByName("id")
	id, err = shared.GetUuidByStrings(inputId)
	if err != nil {
		return id, err
	}
	return id, err
}
