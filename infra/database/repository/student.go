package repository

import (
	"errors"
	"fmt"
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

func (sr *StudentRepository) List() []entities.Student {
	students := memory.GetConnection()
	return students
}

func (sr *StudentRepository) SearchByID(id uuid.UUID) (student entities.Student, err error) {
	students := memory.GetConnection()

	for _, stdu := range students {
		fmt.Println(stdu)
		if stdu.ID == id {
			student = stdu
		}
	}

	fmt.Println(student.ID)
	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Estudante não encontrado aaaaaaaaaaa")
	}

	return student, err
}
