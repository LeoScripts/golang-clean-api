//banco em memoria somente pra testes

package memory

import (
	"golang-student-01/entities"

	"github.com/google/uuid"
)

var StudentsMemory = []*entities.Student{
	&entities.Student{uuid.Must(uuid.Parse("f69aa4db-7bdc-4d07-ad09-9f56ba2a7d6f")), "Alex", 38},
	&entities.Student{uuid.Must(uuid.Parse("98f6c71a-11f4-44b7-8712-32894fea0665")), "Bia", 32},
	&entities.Student{uuid.Must(uuid.Parse("3e35dc08-423b-4485-8783-8b374b1911f3")), "Carlos", 28},
}

func FindAll() []*entities.Student {
	return StudentsMemory
}

func FindOne(id uuid.UUID) (student *entities.Student) {
	for _, sdt := range StudentsMemory {
		if sdt.ID == id {
			student = sdt
		}
	}
	return student
}

func CreatMemory(student *entities.Student) {
	StudentsMemory = append(StudentsMemory, student)
}

func UpdateMemory(student *entities.Student) {
	var newStudents []*entities.Student

	for _, stud := range StudentsMemory {
		if student.ID == stud.ID {
			//atualiza
			newStudents = append(newStudents, student)
		} else {
			// mantem os existentes
			newStudents = append(newStudents, stud)
		}
	}
	StudentsMemory = newStudents
}

func DeleteMemory(id uuid.UUID) {
	var newStudents []*entities.Student
	for _, stdu := range StudentsMemory {
		if id != stdu.ID {
			newStudents = append(newStudents, stdu)
		}
	}
	StudentsMemory = newStudents
}
