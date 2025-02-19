package api

import (
	heart "golang-student-01/api/controller/infra"
	"golang-student-01/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", heart.HeartController)

	groupStudents := s.Engine.Group("/students")
	groupStudents.GET("/", s.StudentController.List)
	groupStudents.POST("/", s.StudentController.Create)
	groupStudents.PUT("/:id", s.StudentController.Update)
	groupStudents.DELETE("/:id", students.Delete)
	groupStudents.GET("/:id", students.Details)
}
