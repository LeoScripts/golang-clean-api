package student

import (
	"github.com/LeoScripts/golang-clean-api/entities"
)

func (su *StudentUsecase) List() (students []*entities.Student) {
	return su.Database.StudentRepository.List()
}
