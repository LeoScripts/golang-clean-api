package student

import (
	"errors"

	"github.com/LeoScripts/golang-clean-api/entities/shared"

	"github.com/google/uuid"
)

func (su *StudentUsecase) Delete(id uuid.UUID) (err error) {
	student := su.Database.StudentRepository.SearchByID(id)
	if student.ID == shared.GetUuidEmpty() {
		return errors.New("Estudante não encontrado")
	}
	su.Database.StudentRepository.Delete(id)
	return err
}
