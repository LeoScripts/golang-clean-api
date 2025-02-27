package student

import (
	"golang-student-01/entities"
)

func (su *StudentUsecase) List() (students []entities.Student) {
	return su.Database.StudentRepository.List()
}
