package api

import (
	"golang-student-01/api/controller/students"
	"golang-student-01/infra/database"
	students_usecase "golang-student-01/usecases/student"

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
