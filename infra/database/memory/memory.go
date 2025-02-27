//banco em memoria somente pra testes

package memory

import (
	"golang-student-01/entities"
	"golang-student-01/entities/shared"
)

func GetConnection() []entities.Student {
	StudentsMock := []entities.Student{
		entities.Student{shared.GetUuid(), "Alex", 38},
		entities.Student{shared.GetUuid(), "Bia", 32},
		entities.Student{shared.GetUuid(), "Carlos", 28},
	}

	return StudentsMock
}
