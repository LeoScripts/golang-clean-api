package student

import (
	"golang-student-01/entities"
)

func (su *StudentUsecase) Create(fullName string, age int) (student *entities.Student, err error) {
	student = entities.NewStudent(fullName, age)
	su.Database.StudentRepository.Create(student)
	return student, err
}
