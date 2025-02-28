package main

import (
	"golang-student-01/api"
	"golang-student-01/infra/database"
	"golang-student-01/infra/database/memory"
	"golang-student-01/infra/database/repository"
)

func main() {
	db := GetDatabase()
	service := api.NewService(db)
	service.Start()
}

func GetDatabase() *database.Database {
	db := memory.StudentsMemory
	studentRepository := repository.NewStudentRepository()
	return database.NewDatabase(db, studentRepository)
}
