package dto

import "time"

type EventDTO struct {
	Name        string
	Description string
	Location    string
	StartDate   time.Time
	EndDate     time.Time
}

type EventsDTO []EventDTO
