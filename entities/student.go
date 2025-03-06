package entities

import (
	"github.com/LeoScripts/golang-clean-api/entities/shared"

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
	SearchByID(id uuid.UUID) *Student
	Create(student *Student)
	Update(student *Student)
	Delete(id uuid.UUID)
}
