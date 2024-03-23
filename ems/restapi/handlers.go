package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/playground/ems/fetcher"
)

func handleHealthCheck() bool {
	return true
}

func handleFetchUserRSVPList(ctx *gin.Context) (string, bool) {
	user, exists := ctx.Get("username")

	return user.(string), exists
}

func handleFetchAllEvents(ctx *gin.Context, app App) (fetcher.EventsData, error) {
	events, err := fetcher.FetchAllEvents(ctx, app.EventFinder)
	if err != nil {
		return fetcher.EventsData{}, err
	}

	return events, nil
}

func handleFetchEventByName(ctx *gin.Context, name string, app App) (fetcher.EventData, error) {
	event, err := fetcher.FetchEventByName(ctx, name, app.EventFinder)
	if err != nil {
		return fetcher.EventData{}, err
	}

	return event, nil
}
