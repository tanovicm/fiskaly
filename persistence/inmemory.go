package persistence

import (
	"fmt"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/google/uuid"
)


type Storage interface {
	CreateDevice(algorithm, label string) (*Device, error)

	// GetDevice retrieves a signature device by its ID
	GetDevice(deviceID string) (*Device, error)

	// ListDevices retrieves all stored signature devices
	ListDevices() ([]Device, error)
}
type InMemoryStorage struct {
	devices map[string]*Device // Map to store signature devices by ID
}

// NewInMemoryStorage creates a new instance of InMemoryStorage
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		devices: make(map[string]*Device),
	}
}
