package entities

import (
	shared "ai/agents/cmd/internal/shared/core/entities"
	"time"
)

type Appointment struct {
	id     *shared.Id
	date   time.Time
	client *Client
}

func (a *Appointment) Id() string {
	return a.id.Value()
}

func (a *Appointment) Date() time.Time {
	return a.date
}

func (a *Appointment) ClientName() string {
	return a.client.Name()
}

func (a *Appointment) ClientGovernmentId() string {
	return a.client.GovernmentId()
}

func NewAppointment(id *shared.Id, date time.Time, client *Client) (*Appointment, error) {
	return &Appointment{id, date, client}, nil
}
