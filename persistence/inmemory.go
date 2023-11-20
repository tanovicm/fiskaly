package persistence

import (
	"fmt"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/google/uuid"
)

type Device struct {
	ID         string // Unique identifier for the device
	Label      string // Label for the device
	PublicKey  []byte // Public key associated with the device
	PrivateKey []byte // Private key associated with the device
	Algorithm  string // Signature algorithm (ECC or RSA)
}

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

// CreateDevice creates a new signature device and stores it in memory
func (s *InMemoryStorage) CreateDevice(algorithm, label string) (*Device, error) {
	publicKeyBytes, privateKeyBytes, err := crypto.GetPairKey(algorithm)
	if err != nil {
		return nil, fmt.Errorf("error creating key pair")
	}
	// Create a new device
	newDevice := &Device{
		ID:         uuid.New().String(),
		Label:      label,
		PublicKey:  publicKeyBytes,
		PrivateKey: privateKeyBytes,
		Algorithm:  algorithm,
	}

	// Store the device in the map
	s.devices[newDevice.ID] = newDevice

	return newDevice, nil
}

// GetDevice retrieves a signature device by its ID from memory
func (s *InMemoryStorage) GetDevice(deviceID string) (*Device, error) {
	device, ok := s.devices[deviceID]
	if !ok {
		return nil, fmt.Errorf("device with id: %s doesn't exist", deviceID)
	}

	return device, nil
}

// ListDevices retrieves all stored signature devices from memory
func (s *InMemoryStorage) ListDevices() ([]Device, error) {
	var deviceList []Device

	for _, device := range s.devices {
		deviceList = append(deviceList, *device)
	}

	return deviceList, nil
}
