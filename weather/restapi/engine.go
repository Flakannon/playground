package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/playground/weather/fetcher"
	"github.com/playground/weather/restapi/middleware"
)

type App struct {
	WeatherClient fetcher.IWeatherDataClient
}

func NewEngine(app App) *gin.Engine {
	r := gin.New()

	weather := r.Group("weather")
	weather.Use(middleware.AuthMiddleWare())
	{
		weather.GET("forecast", getWeatherForecast(app))
		weather.GET("current", getCurrentWeather(app))
	}

	return r
}
