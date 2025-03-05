package student

import (
	"errors"
	"golang-student-01/entities"
	"golang-student-01/entities/shared"
	"golang-student-01/infra/database/memory"

	"github.com/google/uuid"
)

func (su *StudentUsecase) Update(id uuid.UUID, fullname string, age int) (student *entities.Student, err error) {
	studentsMemory := memory.StudentsMemory
	for _, sdt := range studentsMemory {
		if sdt.ID == id {
			student = sdt
		}
	}
	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Estudante não encontrado")
	}

	student.FullName = fullname
	student.Age = age

	su.Database.StudentRepository.Update(student)

	return student, err
}
