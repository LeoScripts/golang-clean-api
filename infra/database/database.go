package database

import (
	"golang-student-01/entities"
)

type Database struct {
	Conn              []*entities.Student
	StudentRepository entities.StudentRespository
}

func NewDatabase(conn []*entities.Student, sr entities.StudentRespository) *Database {
	return &Database{
		Conn:              conn,
		StudentRepository: sr,
	}
}
