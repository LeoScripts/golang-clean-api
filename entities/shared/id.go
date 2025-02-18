package shared

import (
	"log"

	"github.com/google/uuid"
)

func GetUuid() uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Fatal error", err)
	}
	return uuid
}

func GetUuidByStrings(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func GetUuidEmpty() uuid.UUID {
	return uuid.Nil
}
