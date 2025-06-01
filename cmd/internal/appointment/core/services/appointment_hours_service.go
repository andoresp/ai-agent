package services

import "time"

type AppointmentHoursService interface {
	ListNextHours() []time.Time
}
