package handlers

import (
	"github.com/google/uuid"
)

// generateID creates a new unique ID
func generateID() string {
	return uuid.New().String()
}

