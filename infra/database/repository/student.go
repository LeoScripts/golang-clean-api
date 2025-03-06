package repository

import (
	"github.com/LeoScripts/golang-clean-api/entities"
	"github.com/LeoScripts/golang-clean-api/infra/database/memory"

	"github.com/google/uuid"
)

type StudentRepository struct {
	// Ctx context.Context
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		// Ctx: ctx,
	}
}

func (sr *StudentRepository) List() []*entities.Student {
	return memory.FindAll()
}

func (sr *StudentRepository) SearchByID(id uuid.UUID) *entities.Student {
	return memory.FindOne(id)
}

func (sr *StudentRepository) Create(student *entities.Student) {
	memory.Create(student)
}

func (sr *StudentRepository) Update(student *entities.Student) {
	memory.Update(student)
}

func (sr *StudentRepository) Delete(id uuid.UUID) {
	memory.Delete(id)
}
