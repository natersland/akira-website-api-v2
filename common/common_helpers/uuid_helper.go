package commonhelpers

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}
