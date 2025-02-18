package student

import "golang-student-01/entities"

func List() (students []entities.Student, err error) {
	students = entities.StudentsMock

	return students, err
}
