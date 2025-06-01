package entities

import "errors"

type ClientName struct {
	value string
}

func (c *ClientName) Value() string {
	return c.value
}

func NewClientName(value string) (*ClientName, error) {
	if len(value) < 3 {
		return nil, errors.New("ClientName should be at least 3 characters long")
	}

	return &ClientName{value}, nil
}
