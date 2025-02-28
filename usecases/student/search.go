package student

import (
	"golang-student-01/entities"

	"github.com/google/uuid"
)

func (su *StudentUsecase) SearchById(id uuid.UUID) (student *entities.Student, err error) {
	return su.Database.StudentRepository.SearchByID(id)
}
