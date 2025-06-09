package usecases

import (
	"ai/agents/cmd/internal/appointment/core/repositories"
	"ai/agents/cmd/internal/appointment/core/services"
	"slices"
	"time"
)

type ListAvailableAppointmentHoursUseCase struct {
	appointmentRepository   repositories.AppointmentRepository
	appointmentHoursService services.AppointmentHoursService
}

func (uc *ListAvailableAppointmentHoursUseCase) Execute() ([]string, error) {
	nh := uc.appointmentHoursService.ListNextHours()
	aps, err := uc.appointmentRepository.FindAllByDates(nh)

	if err != nil {
		return nil, err
	}

	var apstr []string

	for _, v := range aps {
		apstr = append(apstr, v.Date().Format(time.RFC3339))
	}

	filter := func(yield func(string) bool) {
		for _, h := range nh {
			date := h.Format(time.RFC3339)

			if !slices.Contains(apstr, date) {
				if !yield(date) {
					return
				}
			}
		}
	}

	return slices.Collect(filter), nil
}

func NewListAvailableAppointmentHoursUseCase(repo *repositories.AppointmentRepository, service *services.AppointmentHoursService) *ListAvailableAppointmentHoursUseCase {
	return &ListAvailableAppointmentHoursUseCase{
		appointmentRepository:   *repo,
		appointmentHoursService: *service,
	}
}
