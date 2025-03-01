package api

import (
	heart "golang-student-01/api/controller/infra"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", heart.HeartController)

	groupStudents := s.Engine.Group("/students")
	groupStudents.GET("/", s.StudentController.List)
	groupStudents.POST("/", s.StudentController.Create)
	groupStudents.PUT("/:id", s.StudentController.Update)
	// groupStudents.DELETE("/:id", s.StudentController.Delete)
	groupStudents.GET("/:id", s.StudentController.Details)
}
