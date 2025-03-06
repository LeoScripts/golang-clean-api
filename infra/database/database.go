package database

import (
	"github.com/LeoScripts/golang-clean-api/entities"
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
