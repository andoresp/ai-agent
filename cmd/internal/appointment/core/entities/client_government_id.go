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
	_, err := regexp.Compile("/^[0-9]{3}.[0-9]{3}.[0-9]{3}-[0-9]{2}$/")

	if err != nil {
		return nil, errors.New("invalid client_government_id value")
	}

	return &ClientGovernmentId{value}, nil
}
