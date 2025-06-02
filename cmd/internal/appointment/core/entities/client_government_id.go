package entities

import (
	"errors"
	"regexp"
)

type ClientGovernmentId struct {
	value string
}

func (c *ClientGovernmentId) Value() string {
	return c.value
}

func NewClientGovernmentId(value string) (*ClientGovernmentId, error) {
	r := regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)

	if ok := r.MatchString(value); ok {
		return &ClientGovernmentId{value}, nil
	}

	return nil, errors.New("invalid client_government_id value")
}
