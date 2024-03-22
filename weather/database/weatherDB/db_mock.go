package weatherDB

import (
	"context"
	"time"

	"github.com/playground/weather/database/dto"
)

type WeatherDBClientMock struct{}

func (w WeatherDBClientMock) GetCurrentWeather(ctx context.Context) (dto.CurrentWeatherDTO, error) {
	currentWeather := dto.CurrentWeatherDTO{}

	currentWeather.CurrentWeather.Condition = "dry/sunny"
	currentWeather.CurrentWeather.Tempreature = "20 degrees"
	currentWeather.CurrentWeather.Day = time.Now()

	return currentWeather, nil
}

func (w WeatherDBClientMock) GetWeatherForecast(ctx context.Context) (dto.WeatherForecastDTO, error) {
	currentForecast := dto.WeatherForecastDTO{}

	currentForecast.Forecast = []dto.WeatherDTO{
		{Condition: "dry/sunny", Tempreature: "18 degrees", Day: time.Now().AddDate(0, 0, 1)},
		{Condition: "dry/sunny", Tempreature: "12 degrees", Day: time.Now().AddDate(0, 0, 2)},
		{Condition: "wet/cold", Tempreature: "8 degrees", Day: time.Now().AddDate(0, 0, 3)},
	}

	return currentForecast, nil
}
