package students

import (
	"github.com/gin-gonic/gin"
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
