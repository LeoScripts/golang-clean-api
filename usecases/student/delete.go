package student

import (
	"golang-student-01/entities"

	"github.com/google/uuid"
)

func Delete(id uuid.UUID) (err error) {
	var newStudents []entities.Student

	//verificar se esse usario existe
	// verificar outras questoes
	// o ideal e aplicar softdelete

	for _, stdu := range entities.StudentsMock {
		if stdu.ID != id {
			newStudents = append(newStudents, stdu)
		}
	}
	entities.StudentsMock = newStudents
	return err
}
