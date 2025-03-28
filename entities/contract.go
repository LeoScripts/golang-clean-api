package entities

import "github.com/google/uuid"

type StudentUsecase interface {
	List() []*Student
	SearchByID(id uuid.UUID) (*Student, error)
	Create(fullname string, age int) (*Student, error)
	Update(id uuid.UUID, string, age int) (*Student, error)
	Delete(id uuid.UUID)
}
