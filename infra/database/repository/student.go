package repository

import (
	"errors"
	"golang-student-01/entities"
	"golang-student-01/entities/shared"
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
	students := memory.StudentsMemory
	return students
}

func (sr *StudentRepository) SearchByID(id uuid.UUID) (student *entities.Student, err error) {
	students := memory.StudentsMemory
	for _, stdu := range students {
		if stdu.ID == id {
			student = stdu
		}
	}
	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Estudante não encontrado aaaaaaaaaaa")
	}
	return student, err
}

func (sr *StudentRepository) Create(student *entities.Student) {
	memory.CreatMemory(student)
}

func (sr *StudentRepository) Update(student *entities.Student) {
	memory.UpdateMemory(student)
}
