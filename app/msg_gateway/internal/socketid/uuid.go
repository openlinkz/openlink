package socketid

import "github.com/google/uuid"

var _ Generator = (*uuidGenerator)(nil)

func NewUUIDGenerator() *uuidGenerator {
	return &uuidGenerator{}
}

type uuidGenerator struct{}

func (g *uuidGenerator) NextSid() (string, error) {
	return uuid.New().String(), nil
}
