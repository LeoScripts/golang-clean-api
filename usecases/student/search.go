package student

import (
	"errors"
	"golang-student-01/entities"
	"golang-student-01/entities/shared"

	"github.com/google/uuid"
)

func (su *StudentUsecase) SearchById(id uuid.UUID) (student *entities.Student, err error) {
	student = su.Database.StudentRepository.SearchByID(id)
	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Estudante não encontrado")
	}
	return student, err
}
