package repositories

import (
	"ai/agents/cmd/internal/appointment/core/entities"
	"ai/agents/cmd/internal/appointment/core/repositories"
	shared "ai/agents/cmd/internal/shared/core/entities"
	"errors"
	"slices"
	"time"
)

type AppointmentRepositoryInMemory struct {
	appointments []entities.Appointment
}

// FindAllByDates implements repositories.AppointmentRepository.
func (repo *AppointmentRepositoryInMemory) FindAllByDates(dates []time.Time) (*[]entities.Appointment, error) {
	var datesu []int64

	for _, v := range repo.appointments {
		datesu = append(datesu, v.Date().Unix())
	}

	filter := func(yield func(entities.Appointment) bool) {
		for _, a := range repo.appointments {
			if slices.Contains(datesu, a.Date().Unix()) {
				if !yield(a) {
					return
				}
			}
		}
	}

	aps := slices.Collect(filter)

	return &aps, nil
}

// FindByDate implements repositories.AppointmentRepository.
func (repo *AppointmentRepositoryInMemory) FindByDate(date time.Time) (*entities.Appointment, error) {
	idx := slices.IndexFunc(repo.appointments, func(a entities.Appointment) bool {
		return a.Date().Unix() == date.Unix()
	})

	if idx >= 0 {
		return &repo.appointments[idx], nil
	}

	return nil, errors.New("appointment not found")

}

// FindByGovernmentId implements repositories.AppointmentRepository.
func (repo *AppointmentRepositoryInMemory) FindByGovernmentId(governmentId entities.ClientGovernmentId) (*entities.Appointment, error) {
	idx := slices.IndexFunc(repo.appointments, func(a entities.Appointment) bool {
		return a.ClientGovernmentId() == governmentId.Value()
	})

	if idx >= 0 {
		return &repo.appointments[idx], nil
	}

	return nil, errors.New("appointment not found")
}

// FindByClientName implements repositories.AppointmentRepository.
func (repo *AppointmentRepositoryInMemory) FindByClientName(name entities.ClientName) (*entities.Appointment, error) {
	idx := slices.IndexFunc(repo.appointments, func(a entities.Appointment) bool {
		return a.ClientName() == name.Value()
	})

	if idx >= 0 {
		return &repo.appointments[idx], nil
	}

	return nil, errors.New("appointment not found")
}

// FindById implements repositories.AppointmentRepository.
func (repo *AppointmentRepositoryInMemory) FindById(id shared.Id) (*entities.Appointment, error) {
	idx := slices.IndexFunc(repo.appointments, func(a entities.Appointment) bool {
		return a.Id() == id.Value()
	})

	if idx >= 0 {
		return &repo.appointments[idx], nil
	}

	return nil, errors.New("appointment not found")
}

// Create implements repositories.AppointmentRepository.
func (repo *AppointmentRepositoryInMemory) Create(date time.Time, client entities.Client) (*entities.Appointment, error) {
	id, err := shared.NewId()

	if err != nil {
		return nil, err
	}

	a, err := entities.NewAppointment(*id, date, client)

	if err != nil {
		return nil, err
	}

	repo.appointments = append(repo.appointments, *a)

	return a, nil
}

// Delete implements repositories.AppointmentRepository.
func (repo *AppointmentRepositoryInMemory) Delete(appointment entities.Appointment) error {
	idx := -1

	for i, v := range repo.appointments {
		if v.Id() == appointment.Id() {
			idx = i
		}
	}

	if idx >= 0 {
		repo.appointments = slices.Delete(repo.appointments, idx, 1)
		return nil
	}

	return errors.New("failed to delete appointment")
}

func NewAppointmentRepositoryInMemory() repositories.AppointmentRepository {
	return &AppointmentRepositoryInMemory{}
}
