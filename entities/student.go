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

// mock
var StudentsMock = []Student{
	// substituimos o id pelo uuid
	Student{shared.GetUuid(), "Alex", 38},
	Student{shared.GetUuid(), "Bia", 32},
	Student{shared.GetUuid(), "Carlos", 28},
}
