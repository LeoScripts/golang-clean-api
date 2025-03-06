package api

import (
	heart "github.com/LeoScripts/golang-clean-api/api/controller/infra"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/api/v1/heart", heart.HeartController)

	groupStudents := s.Engine.Group("/api/v1/students")
	groupStudents.GET("/", s.StudentController.List)
	groupStudents.POST("/", s.StudentController.Create)
	groupStudents.PUT("/:id", s.StudentController.Update)
	groupStudents.DELETE("/:id", s.StudentController.Delete)
	groupStudents.GET("/:id", s.StudentController.Details)
}
