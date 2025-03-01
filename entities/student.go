package entities

import (
	"golang-student-01/entities/shared"

	"github.com/google/uuid"
)

type Student struct {
	ID       uuid.UUID `json:"_id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}

func NewStudent(fullName string, age int) *Student {
	return &Student{
		ID:       shared.GetUuid(),
		FullName: fullName,
		Age:      age,
	}
}

type StudentRespository interface {
	List() []*Student
	SearchByID(id uuid.UUID) (*Student, error)
	Create(student *Student)
	Update(student *Student)
}
