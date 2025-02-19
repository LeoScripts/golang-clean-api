package api

import (
	"golang-student-01/api/controller/students"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Engine            *gin.Engine
	StudentController *students.StudentController
}

func NewService() *Service {
	return &Service{
		Engine: gin.Default(),
	}
}

func (s *Service) Start() {
	s.GetRoutes()
	s.Engine.Run(":7777")
}
