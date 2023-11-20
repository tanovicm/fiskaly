package crypto

import "sync"

type SignerImpl struct {
	signatureCounter int
	lastSignature    string
	mu               sync.Mutex // Mutex to ensure safe concurrent access
}

func NewSigner() *SignerImpl {
	return &SignerImpl{signatureCounter: 0, lastSignature: ""}
}

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
	IncrementSignatureCounter()
	GetSignatureCounter() int
}

func (s *SignerImpl) Sign(data []byte) ([]byte, error) {
	if s.GetSignatureCounter() == 0 {
		return nil, nil
	}
	s.IncrementSignatureCounter()
	return nil, nil
}

// IncrementSignatureCounter increments the counter by 1
func (s *SignerImpl) IncrementSignatureCounter() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.signatureCounter++
}

// GetSignatureCounter returns the current value of the counter
func (s *SignerImpl) GetSignatureCounter() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.signatureCounter
}

// // IncrementSignatureCounter increments the counter by 1
// func (s *SignerImpl) SetLastSignature(lastSignature string) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	s.lastSignature = lastSignature
// }

// // GetSignatureCounter returns the current value of the counter
// func (s *SignerImpl) GetLastSignature() string {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	return s.lastSignature
// }
