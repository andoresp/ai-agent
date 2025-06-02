package services

import (
	"ai/agents/cmd/internal/appointment/core/services"
	"time"
)

type AppointmentHoursServiceImpl struct {
}

func (as *AppointmentHoursServiceImpl) ListNextHours() []time.Time {
	var dates []time.Time
	dates = append(dates, time.Now())
	return dates
}

func NewAppointmentHoursService() services.AppointmentHoursService {
	return &AppointmentHoursServiceImpl{}
}
