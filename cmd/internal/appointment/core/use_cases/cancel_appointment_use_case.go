package usecases

import (
	"ai/agents/cmd/internal/appointment/core/repositories"
	shared "ai/agents/cmd/internal/shared/core/entities"
)

type CancelAppointmentDto struct {
	Id string
}

type CancelAppointmentUseCase struct {
	appointmentRepository repositories.AppointmentRepository
}

func (uc *CancelAppointmentUseCase) Execute(dto CancelAppointmentDto) error {
	id, err := shared.NewId()

	if err != nil {
		return err
	}

	a, err := uc.appointmentRepository.FindById(*id)

	if err != nil {
		return err
	}

	err = uc.appointmentRepository.Delete(*a)

	if err != nil {
		return err
	}

	return nil
}
