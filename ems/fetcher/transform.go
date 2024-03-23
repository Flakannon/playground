package fetcher

import "github.com/playground/ems/database/dto"

type EventData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartDate   string `json:"start"`
	EndDate     string `json:"end"`
}

func toEventData(rawEventData dto.EventDTO) EventData {
	var eventInfo EventData

	eventInfo.Name = rawEventData.Name
	eventInfo.Description = rawEventData.Description
	eventInfo.Location = rawEventData.Location

	startDateString := rawEventData.StartDate.Format("2006-01-02 15:04:05")
	eventInfo.StartDate = startDateString

	endDateString := rawEventData.EndDate.Format("2006-01-02 15:04:05")
	eventInfo.EndDate = endDateString

	return eventInfo
}

type EventsData []EventData

func toEventsData(rawEventsData dto.EventsDTO) EventsData {
	var events EventsData

	for _, event := range rawEventsData {
		eventInfo := toEventData(event)
		events = append(events, eventInfo)
	}

	return events
}
