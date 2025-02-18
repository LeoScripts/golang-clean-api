package api

import (
	heart "golang-student-01/api/controller/infra"
	"golang-student-01/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", heart.HeartController)

	groupStudents := s.Group("/students")
	groupStudents.GET("/", students.List)
	groupStudents.POST("/", students.Create)
	groupStudents.PUT("/:id", students.Update)
	groupStudents.DELETE("/:id", students.Delete)
	groupStudents.GET("/:id", students.Details)
}
