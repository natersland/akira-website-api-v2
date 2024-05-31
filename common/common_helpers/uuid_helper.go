package commonhelpers

import (
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

func GetCurrentTimeISO() string {
	currentTime := time.Now().UTC()         // Get the current UTC time
	return currentTime.Format(time.RFC3339) // Format as ISO 8601
}
