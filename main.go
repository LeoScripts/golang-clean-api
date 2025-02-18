package main

import (
	"golang-student-01/api"
)

func main() {
	service := api.NewService()

	service.Start()
}
