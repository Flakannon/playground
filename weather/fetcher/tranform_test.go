package fetcher

import (
	"reflect"
	"testing"
	"time"

	"github.com/playground/weather/database/dto"
)

func Test_toForecastWeatherData(t *testing.T) {
	type args struct {
		forecast dto.WeatherForecastDTO
	}
	tests := []struct {
		name string
		args args
		want WeatherForecast
	}{
		{
			"test transform forecast",
			args{
				forecast: dto.WeatherForecastDTO{
					Forecast: []dto.WeatherDTO{
						{Condition: "dry/sunny", Tempreature: "18 degrees", Day: time.Now().AddDate(0, 0, 1)},
						{Condition: "dry/sunny", Tempreature: "12 degrees", Day: time.Now().AddDate(0, 0, 2)},
						{Condition: "wet/cold", Tempreature: "8 degrees", Day: time.Now().AddDate(0, 0, 3)},
					},
				},
			},
			WeatherForecast{
				ForecastReport: []WeatherReport{
					{Condition: "dry/sunny", Temp: "18 degrees", Date: time.Now().AddDate(0, 0, 1).Format("2006-01-02")},
					{Condition: "dry/sunny", Temp: "12 degrees", Date: time.Now().AddDate(0, 0, 2).Format("2006-01-02")},
					{Condition: "wet/cold", Temp: "8 degrees", Date: time.Now().AddDate(0, 0, 3).Format("2006-01-02")},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toForecastWeatherData(tt.args.forecast); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toForecastWeatherData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toCurrentWeatherData(t *testing.T) {
	type args struct {
		weatherData dto.CurrentWeatherDTO
	}
	tests := []struct {
		name string
		args args
		want CurrentWeather
	}{
		{
			"test transform current weather", args{
				weatherData: dto.CurrentWeatherDTO{
					CurrentWeather: dto.WeatherDTO{
						Condition:   "dry",
						Tempreature: "20 degrees",
						Day:         time.Now(),
					},
				},
			}, CurrentWeather{
				WeatherReport: WeatherReport{
					Condition: "dry",
					Temp:      "20 degrees",
					Date:      time.Now().Format("2006-01-02"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toCurrentWeatherData(tt.args.weatherData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toCurrentWeatherData() = %v, want %v", got, tt.want)
			}
		})
	}
}
