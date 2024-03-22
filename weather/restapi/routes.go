package restapi

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCurrentWeather(app App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := handleCurrentWeatherRequest(app)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		bytes, err := json.Marshal(data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		ctx.Data(http.StatusOK, "application/json", bytes)
	}
}

func getWeatherForecast(app App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := handleWeatherForeCastRequest(app)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		bytes, err := json.Marshal(data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		ctx.Data(http.StatusOK, "application/json", bytes)
	}
}
