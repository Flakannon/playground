package fetcher

import (
	"context"

	"github.com/playground/weather/database/dto"
)

type IWeatherDataClient interface {
	GetCurrentWeather(ctx context.Context) (dto.CurrentWeatherDTO, error)
	GetWeatherForecast(ctx context.Context) (dto.WeatherForecastDTO, error)
}

func FetchWeatherForecast(ctx context.Context, weatherClient IWeatherDataClient) (WeatherForecast, error) {
	weatherData, err := weatherClient.GetWeatherForecast(ctx)
	if err != nil {
		return WeatherForecast{}, err
	}

	transformed := toForecastWeatherData(weatherData)

	return transformed, nil
}

func FetchCurrentWeather(ctx context.Context, weatherClient IWeatherDataClient) (CurrentWeather, error) {
	weatherData, err := weatherClient.GetCurrentWeather(ctx)
	if err != nil {
		return CurrentWeather{}, err
	}

	transformed := toCurrentWeatherData(weatherData)

	return transformed, nil
}
