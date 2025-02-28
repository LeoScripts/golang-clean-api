//banco em memoria somente pra testes

package memory

import (
	"golang-student-01/entities"

	"github.com/google/uuid"
)

func GetConnection() []entities.Student {
	StudentsMock := []entities.Student{
		entities.Student{uuid.Must(uuid.Parse("f69aa4db-7bdc-4d07-ad09-9f56ba2a7d6f")), "Alex", 38},
		entities.Student{uuid.Must(uuid.Parse("98f6c71a-11f4-44b7-8712-32894fea0665")), "Bia", 32},
		entities.Student{uuid.Must(uuid.Parse("3e35dc08-423b-4485-8783-8b374b1911f3")), "Carlos", 28},
	}

	return StudentsMock
}
