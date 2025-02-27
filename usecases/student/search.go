package student

import (
	"golang-student-01/entities"

	"github.com/google/uuid"
)

func (su *StudentUsecase) SearchById(id uuid.UUID) (student entities.Student, err error) {
	// StudentsMock := memory.GetConnection()

	// for _, stdu := range StudentsMock {
	// 	if id == stdu.ID {
	// 		student = stdu
	// 	}
	// }

	// student = su.repository.Details(id)

	// if student.ID == shared.GetUuidEmpty() {
	// 	return student, errors.New("Estudante não encontrado aaaaaaaaaaa")
	// }

	return student, err
}
