package api

import (
	"github.com/LeoScripts/golang-clean-api/api/controller/students"
	"github.com/LeoScripts/golang-clean-api/infra/database"
	students_usecase "github.com/LeoScripts/golang-clean-api/usecases/student"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Engine            *gin.Engine
	DB                *database.Database
	StudentController *students.StudentController
}

func NewService(db *database.Database) *Service {
	return &Service{
		Engine: gin.Default(),
		DB:     db,
	}
}

func (s *Service) Start() {
	s.GetControllers()
	s.GetRoutes()
	s.Engine.Run(":7777")
}

func (s *Service) GetControllers() {
	studentUsecase := students_usecase.NewStudentUsecase(s.DB)
	s.StudentController = students.NewStudentController(studentUsecase)
}
