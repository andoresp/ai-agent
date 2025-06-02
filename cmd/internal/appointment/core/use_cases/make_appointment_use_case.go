package usecases

import (
	"ai/agents/cmd/internal/appointment/core/entities"
	"ai/agents/cmd/internal/appointment/core/repositories"
	"ai/agents/cmd/internal/appointment/core/services"
	"time"
)

type MakeAppointmentDto struct {
	Date               string `json:"date"`
	ClientName         string `json:"client_name"`
	ClientGovernmentId string `json:"client_government_id"`
}

type MakeAppointmentUseCase struct {
	appointmentRepository   repositories.AppointmentRepository
	appointmentHoursService services.AppointmentHoursService
}

func (uc *MakeAppointmentUseCase) Execute(dto MakeAppointmentDto) (*entities.Appointment, error) {

	cn, err := entities.NewClientName(dto.ClientName)

	if err != nil {
		return nil, err
	}

	cgi, err := entities.NewClientGovernmentId(dto.ClientGovernmentId)

	if err != nil {
		return nil, err
	}

	c, err := entities.NewClient(*cn, *cgi)

	if err != nil {
		return nil, err
	}

	a, err := uc.appointmentRepository.Create(time.Now(), *c)

	if err != nil {
		return nil, err
	}

	return a, nil
}

func NewMakeAppoinmentUseCase(repo *repositories.AppointmentRepository, service *services.AppointmentHoursService) *MakeAppointmentUseCase {
	return &MakeAppointmentUseCase{
		appointmentRepository:   *repo,
		appointmentHoursService: *service,
	}
}
