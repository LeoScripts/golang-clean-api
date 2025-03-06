package student

import (
	"github.com/LeoScripts/golang-clean-api/entities"
)

func (su *StudentUsecase) Create(fullName string, age int) (student *entities.Student, err error) {
	student = entities.NewStudent(fullName, age)
	su.Database.StudentRepository.Create(student)
	return student, err
}
