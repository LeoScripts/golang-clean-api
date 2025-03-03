package repository

import (
	"golang-student-01/entities"
	"golang-student-01/infra/database/memory"

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

func (sr *StudentRepository) SearchByID(id uuid.UUID) (student *entities.Student) {
	return memory.FindOne(id)
}

func (sr *StudentRepository) Create(student *entities.Student) {
	memory.CreatMemory(student)
}

func (sr *StudentRepository) Update(student *entities.Student) {
	memory.UpdateMemory(student)
}

func (sr *StudentRepository) Delete(id uuid.UUID) {
	memory.DeleteMemory(id)
}
