package fetcher

import (
	"reflect"
	"testing"
	"time"

	"github.com/playground/ems/database/dto"
)

func Test_toEventData(t *testing.T) {
	type args struct {
		rawEventData dto.EventDTO
	}
	tests := []struct {
		name string
		args args
		want EventData
	}{
		{"transformEvent", args{rawEventData: dto.EventDTO{
			Name:        "testName",
			Description: "testDesc",
			Location:    "testLocation",
			StartDate:   time.Date(2025, 1, 1, 1, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 1, 1, 1, 0, 0, 0, time.UTC),
		}}, EventData{
			Name:        "testName",
			Description: "testDesc",
			Location:    "testLocation",
			StartDate:   "2025-01-01 01:00:00",
			EndDate:     "2025-01-01 01:00:00",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toEventData(tt.args.rawEventData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toEventData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toEventsData(t *testing.T) {
	type args struct {
		rawEventsData dto.EventsDTO
	}
	tests := []struct {
		name string
		args args
		want EventsData
	}{
		{"transformEvent", args{rawEventsData: dto.EventsDTO{
			{
				Name:        "testName",
				Description: "testDesc",
				Location:    "testLocation",
				StartDate:   time.Date(2025, 1, 1, 1, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2025, 1, 1, 1, 0, 0, 0, time.UTC),
			},
			{
				Name:        "testName2",
				Description: "testDesc2",
				Location:    "testLocation2",
				StartDate:   time.Date(2026, 1, 1, 1, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2026, 1, 1, 1, 0, 0, 0, time.UTC),
			},
		}}, EventsData{
			{
				Name:        "testName",
				Description: "testDesc",
				Location:    "testLocation",
				StartDate:   "2025-01-01 01:00:00",
				EndDate:     "2025-01-01 01:00:00",
			},
			{
				Name:        "testName2",
				Description: "testDesc2",
				Location:    "testLocation2",
				StartDate:   "2026-01-01 01:00:00",
				EndDate:     "2026-01-01 01:00:00",
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toEventsData(tt.args.rawEventsData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toEventsData() = %v, want %v", got, tt.want)
			}
		})
	}
}
