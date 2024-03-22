package weatherDB

import (
	"context"

	"github.com/playground/weather/database"
	"github.com/playground/weather/database/dto"
)

type WeatherDBClient struct {
	database.BaseClient
}

func (w WeatherDBClient) GetCurrentWeather(ctx context.Context) (dto.CurrentWeatherDTO, error) {
	return dto.CurrentWeatherDTO{}, nil
}

func (w WeatherDBClient) GetWeatherForecast(ctx context.Context) (dto.WeatherForecastDTO, error) {
	return dto.WeatherForecastDTO{}, nil
}
