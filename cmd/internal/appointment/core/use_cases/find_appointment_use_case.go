package usecases

import (
	"ai/agents/cmd/internal/appointment/core/entities"
	"ai/agents/cmd/internal/appointment/core/repositories"
	"errors"
)

type FindAppointmentDto struct {
	ClientName           string
	ClientGovernmentName string
}

type FindAppointmentUseCase struct {
	appointmentRepository repositories.AppointmentRepository
}

func (uc *FindAppointmentUseCase) Execute(dto FindAppointmentDto) (*entities.Appointment, error) {
	if dto.ClientGovernmentName != "" {
		cgn, err := entities.NewClientGovernmentId(dto.ClientGovernmentName)

		if err != nil {
			return nil, err
		}

		a, err := uc.appointmentRepository.FindByGovernmentId(cgn)

		if err != nil {
			return nil, err
		}

		return a, nil
	}

	if dto.ClientName != "" {
		cn, err := entities.NewClientName(dto.ClientName)

		if err != nil {
			return nil, err
		}

		a, err := uc.appointmentRepository.FindByClientName(cn)

		if err != nil {
			return nil, err
		}

		return a, nil
	}

	return nil, errors.New("appointment not found")
}

func NewFindAppointmentUseCase(repo *repositories.AppointmentRepository) *FindAppointmentUseCase {
	return &FindAppointmentUseCase{
		appointmentRepository: *repo,
	}
}
