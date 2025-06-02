package entities

import (
	"errors"
	"math/rand"
)

type Id struct {
	value string
}

func (i *Id) Value() string {
	return i.value
}

func NewId(id ...string) (*Id, error) {
	size := 20

	if len(id) > 0 {
		if len(id[0]) < size {
			return nil, errors.New("invalid id size")
		}

		return &Id{id[0]}, nil
	}

	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var chars []byte

	for i := 0; i < size; i += 1 {
		index := rand.Intn(len(alphabet))
		char := alphabet[index]
		chars = append(chars, char)
	}

	return &Id{string(chars)}, nil
}
