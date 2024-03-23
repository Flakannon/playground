package fetcher

import (
	"context"

	"github.com/playground/ems/database/dto"
)

type IEventFinder interface {
	GetEvents() (dto.EventsDTO, error)
	GetEventByName(string) (dto.EventDTO, error)
}

func FetchEventByName(ctx context.Context, name string, eventFinder IEventFinder) (EventData, error) {
	rawInfo, err := eventFinder.GetEventByName(name)
	if err != nil {
		return EventData{}, err
	}
	return toEventData(rawInfo), nil
}

func FetchAllEvents(ctx context.Context, eventFinder IEventFinder) (EventsData, error) {
	rawInfo, err := eventFinder.GetEvents()
	if err != nil {
		return EventsData{}, err
	}
	return toEventsData(rawInfo), nil
}
