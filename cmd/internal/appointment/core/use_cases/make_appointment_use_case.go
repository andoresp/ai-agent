package usecases

import (
	"ai/agents/cmd/internal/appointment/core/entities"
	"ai/agents/cmd/internal/appointment/core/repositories"
	"ai/agents/cmd/internal/appointment/core/services"
	"errors"
	"fmt"
	"slices"
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

	t, err := time.Parse(time.RFC3339, dto.Date)

	if err != nil {
		return nil, fmt.Errorf("invalid time %s", err.Error())
	}

	if time.Now().After(t.Add(time.Duration(time.Second))) {
		return nil, errors.New("unable to make appointment, because, the given date already past")
	}

	cn, err := entities.NewClientName(dto.ClientName)

	if err != nil {
		return nil, err
	}

	cgi, err := entities.NewClientGovernmentId(dto.ClientGovernmentId)

	if err != nil {
		return nil, err
	}

	if ok, _ := uc.appointmentRepository.FindByGovernmentId(cgi); ok != nil {
		return nil, errors.New("unable to make appointment, because, an appointment for the given client already exists")
	}

	if ok, _ := uc.appointmentRepository.FindByDate(t); ok != nil {
		return nil, errors.New("unable to make appointment, because, the date is already taken")
	}

	c := entities.NewClient(cn, cgi)
	naps := uc.appointmentHoursService.ListNextHours()

	isTimeAvailable := slices.ContainsFunc(naps, func(nat time.Time) bool {
		return nat.Format(time.RFC3339) == t.Format(time.RFC3339)
	})

	if !isTimeAvailable {
		return nil, errors.New("unable to make appointment, because, the date is not available")
	}

	a, err := uc.appointmentRepository.Create(t, c)

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
