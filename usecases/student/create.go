package student

import (
	"golang-student-01/entities"
)

func Create(fullName string, age int) (student entities.Student, err error) {

	// pointStudent := entities.NewStudent(fullName, age)
	// student = *pointStudent
	// entities.StudentsMock = append(entities.StudentsMock, student)

	return student, err
}
