package student

import (
	"errors"

	"golang-student-01/entities"
	"golang-student-01/entities/shared"

	"github.com/google/uuid"
)

func SearchById(id uuid.UUID) (student entities.Student, err error) {
	for _, stdu := range entities.StudentsMock {
		if id == stdu.ID {
			student = stdu
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Estudante não encontrado")
	}

	return student, err
}
