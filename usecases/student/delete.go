package student

import (
	"errors"

	"golang-student-01/entities/shared"

	"github.com/google/uuid"
)

func (su *StudentUsecase) Delete(id uuid.UUID) (err error) {
	student, err := su.Database.StudentRepository.SearchByID(id)
	if err != nil {
		return err
	}
	if student.ID == shared.GetUuidEmpty() {
		return errors.New("id nao encotrado")
	}
	su.Database.StudentRepository.Delete(id)
	return err
}
