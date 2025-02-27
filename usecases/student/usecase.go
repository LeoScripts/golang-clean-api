package student

import (
	"golang-student-01/infra/database"
)

type StudentUsecase struct {
	Database *database.Database
}

func NewStudentUsecase(db *database.Database) *StudentUsecase {
	return &StudentUsecase{
		Database: db,
	}
}
