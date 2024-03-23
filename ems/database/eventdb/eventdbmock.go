package eventdb

import (
	"errors"
	"time"

	"github.com/playground/ems/database/dto"
)

type EventDBMock struct{}

func (e EventDBMock) GetEvents() (dto.EventsDTO, error) {
	var events dto.EventsDTO

	event1 := dto.EventDTO{
		Name:        "GoConf",
		Description: "Go tips and tricks with slightly aggressive german accents",
		Location:    "Berlin",
		StartDate:   time.Date(2025, 2, 12, 12, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2025, 2, 12, 18, 0, 0, 0, time.UTC),
	}

	events = append(events, event1)

	event2 := dto.EventDTO{
		Name:        "PokemonConf",
		Description: "Nerds unite at pokeconf",
		Location:    "Tokyo",
		StartDate:   time.Date(2027, 2, 14, 12, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2027, 2, 12, 18, 0, 0, 0, time.UTC),
	}

	events = append(events, event2)

	return events, nil
}

func (e EventDBMock) GetEventByName(nameSearched string) (dto.EventDTO, error) {
	var events dto.EventsDTO

	event1 := dto.EventDTO{
		Name:        "GoConf",
		Description: "Go tips and tricks with slightly aggressive german accents",
		Location:    "Berlin",
		StartDate:   time.Date(2025, 2, 12, 12, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2025, 2, 12, 18, 0, 0, 0, time.UTC),
	}

	events = append(events, event1)

	event2 := dto.EventDTO{
		Name:        "PokemonConf",
		Description: "Nerds unite at pokeconf",
		Location:    "Tokyo",
		StartDate:   time.Date(2027, 2, 14, 12, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2027, 2, 12, 18, 0, 0, 0, time.UTC),
	}

	events = append(events, event2)

	for _, event := range events {
		if event.Name == nameSearched {
			return event, nil
		}
	}

	return dto.EventDTO{}, errors.New("event not found")
}
