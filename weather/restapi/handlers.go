package restapi

import (
	"context"

	"github.com/playground/weather/fetcher"
)

func handleCurrentWeatherRequest(app App) (fetcher.CurrentWeather, error) {
	// call nusiness layer and rtetuinr data and error
	data, err := fetcher.FetchCurrentWeather(context.Background(), app.WeatherClient)
	if err != nil {
		return fetcher.CurrentWeather{}, err
	}
	return data, nil
}

func handleWeatherForeCastRequest(app App) (fetcher.WeatherForecast, error) {
	// call business layer and return data and error
	data, err := fetcher.FetchWeatherForecast(context.Background(), app.WeatherClient)
	if err != nil {
		return fetcher.WeatherForecast{}, err
	}
	return data, nil
}
