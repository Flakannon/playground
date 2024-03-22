package fetcher

import (
	"github.com/playground/weather/database/dto"
)

type WeatherReport struct {
	Condition string `json:"condition"`
	Temp      string `json:"temp"`
	Date      string `json:"date"`
}

type CurrentWeather struct {
	WeatherReport
}

func toCurrentWeatherData(weatherData dto.CurrentWeatherDTO) CurrentWeather {
	var data CurrentWeather
	data.Condition = weatherData.CurrentWeather.Condition
	data.Temp = weatherData.CurrentWeather.Tempreature
	data.Date = weatherData.CurrentWeather.Day.Format("2006-01-02")

	return data
}

type ForecastReport []WeatherReport

type WeatherForecast struct {
	ForecastReport
}

func toForecastWeatherData(forecast dto.WeatherForecastDTO) WeatherForecast {
	var data WeatherForecast

	for _, weatherForDay := range forecast.Forecast {
		var report WeatherReport

		report.Condition = weatherForDay.Condition
		report.Temp = weatherForDay.Tempreature
		report.Date = weatherForDay.Day.Format("2006-01-02")

		data.ForecastReport = append(data.ForecastReport, report)
	}

	return data
}
