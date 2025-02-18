package student

import (
	"errors"
	"golang-student-01/entities"
	"golang-student-01/entities/shared"

	"github.com/google/uuid"
)

func Update(id uuid.UUID, fullname string, age int) (student entities.Student, err error) {
	var newStudents []entities.Student

	for _, sdt := range entities.StudentsMock {
		if sdt.ID == id {
			student = sdt
		}
	}
	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("id nao encotrado")
	}

	student.FullName = fullname
	student.Age = age

	for _, stud := range entities.StudentsMock {
		if student.ID == stud.ID {
			newStudents = append(newStudents, student)
		} else {
			newStudents = append(newStudents, stud)
		}
	}

	entities.StudentsMock = newStudents
	return student, err
}
