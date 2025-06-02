package usecases

import (
	"ai/agents/cmd/internal/appointment/core/entities"
	"ai/agents/cmd/internal/appointment/core/repositories"
	"ai/agents/cmd/internal/appointment/core/services"
)

type ListAvailableAppointmentHoursUseCase struct {
	appointmentRepository   repositories.AppointmentRepository
	appointmentHoursService services.AppointmentHoursService
}

func (uc *ListAvailableAppointmentHoursUseCase) Execute() (*[]entities.Appointment, error) {
	nh := uc.appointmentHoursService.ListNextHours()
	ah, err := uc.appointmentRepository.FindAllByDates(nh)

	if err != nil {
		return nil, err
	}
	return ah, nil
}

func NewListAvailableAppointmentHoursUseCase(repo *repositories.AppointmentRepository, service *services.AppointmentHoursService) *ListAvailableAppointmentHoursUseCase {
	return &ListAvailableAppointmentHoursUseCase{
		appointmentRepository:   *repo,
		appointmentHoursService: *service,
	}
}
