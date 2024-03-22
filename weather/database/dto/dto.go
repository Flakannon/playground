package dto

import "time"

type WeatherDTO struct {
	Condition   string
	Tempreature string
	Day         time.Time
}

type CurrentWeatherDTO struct {
	CurrentWeather WeatherDTO
}

type WeatherForecastDTO struct {
	Forecast []WeatherDTO
}
