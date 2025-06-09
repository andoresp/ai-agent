package services

import (
	"ai/agents/cmd/internal/appointment/core/services"
	"time"
)

type AppointmentHoursServiceImpl struct {
	daysAhead    int
	openingHour  int
	workingHours []int
}

var defaultWorkingHours = []int{8, 9, 10, 11, 13, 14, 15, 16}

func (as *AppointmentHoursServiceImpl) ListNextHours() []time.Time {
	var dates []time.Time

	date := time.Now()

	for i := 0; i < as.daysAhead; i += 1 {
		date := time.Date(
			date.Year(),
			date.Month(),
			date.Day()+i,
			0, 0, 0, 0,
			date.Location(),
		)

		// add working hours to dates
		for _, j := range as.workingHours {
			date := time.Date(
				date.Year(),
				date.Month(),
				date.Day(),
				j, 0, 0, 0,
				date.Location(),
			)

			if date.Before(time.Now()) {
				continue
			}

			dates = append(dates, date)
		}
	}

	return dates
}

func NewAppointmentHoursService(daysAhead int, openingHour int, workingHours ...int) services.AppointmentHoursService {
	if workingHours == nil {
		workingHours = defaultWorkingHours
	}

	return &AppointmentHoursServiceImpl{daysAhead, openingHour, workingHours}
}
