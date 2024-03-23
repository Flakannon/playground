package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHealth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		health := handleHealthCheck()
		if !health {
			ctx.JSON(http.StatusInternalServerError, "Not healthy")
			return
		}

		ctx.JSON(http.StatusOK, "Healthy")
	}
}

func getRSVPListForUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, exists := handleFetchUserRSVPList(ctx)
		if !exists {
			ctx.JSON(http.StatusBadRequest, "User has not been found")
			return
		}
		ctx.JSON(http.StatusOK, fmt.Sprintf("Returned RSVP list for %v", user))
	}
}

func getAllEventsInfo(app App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		events, err := handleFetchAllEvents(ctx, app)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		data, err := json.Marshal(events)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Data(http.StatusOK, "application/json", data)
	}
}

func getEventByName(app App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eventName := ctx.Param("name")
		events, err := handleFetchEventByName(ctx, eventName, app)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		data, err := json.Marshal(events)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Data(http.StatusOK, "application/json", data)
	}
}
