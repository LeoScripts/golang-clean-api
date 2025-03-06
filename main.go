package main

import (
	"github.com/LeoScripts/golang-clean-api/api"
	"github.com/LeoScripts/golang-clean-api/infra/database"
	"github.com/LeoScripts/golang-clean-api/infra/database/memory"
	"github.com/LeoScripts/golang-clean-api/infra/database/repository"
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
