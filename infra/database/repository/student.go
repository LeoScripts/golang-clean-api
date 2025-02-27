package repository

import (
	"golang-student-01/entities"
	"golang-student-01/infra/database/memory"
)

type StudentRepository struct {
	// Ctx context.Context
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		// Ctx: ctx,
	}
}

func (sr *StudentRepository) List() []entities.Student {
	students := memory.GetConnection()
	return students
}
