package repositories

import (
	"ai/agents/cmd/internal/appointment/core/entities"
	shared "ai/agents/cmd/internal/shared/core/entities"
	"time"
)

type AppointmentRepository interface {
	FindById(id shared.Id) (*entities.Appointment, error)
	FindByClientName(name entities.ClientName) (*entities.Appointment, error)
	FindByGovernmentId(governmentId entities.ClientGovernmentId) (*entities.Appointment, error)
	FindByDate(date time.Time) (*entities.Appointment, error)
	FindAllByDates(dates []time.Time) ([]*entities.Appointment, error)
	Create(date time.Time, client entities.Client) (*entities.Appointment, error)
	Delete(appointment entities.Appointment) error
}
