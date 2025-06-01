package entities

import "math/rand"

type Id struct {
	value string
}

func (i *Id) Value() string {
	return i.value
}

func NewId() (*Id, error) {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	size := 20

	var chars []byte

	for i := 0; i < size; i += 1 {
		index := rand.Intn(len(alphabet))
		char := alphabet[index]
		chars = append(chars, char)
	}

	return &Id{string(chars)}, nil
}
