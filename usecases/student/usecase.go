package student

import (
	"github.com/LeoScripts/golang-clean-api/infra/database"
)

type StudentUsecase struct {
	Database *database.Database
}

func NewStudentUsecase(db *database.Database) *StudentUsecase {
	return &StudentUsecase{
		Database: db,
	}
}
