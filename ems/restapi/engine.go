package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/playground/ems/fetcher"
	"github.com/playground/ems/restapi/middleware"
)

type App struct {
	EventFinder fetcher.IEventFinder
}

func NewEngine(app App) *gin.Engine {
	r := gin.New()

	r.GET("/health", getHealth())

	eventInfo := r.Group("info")
	{
		eventInfo.GET("/allevents", getAllEventsInfo(app))
		eventInfo.GET("/event/:name", getEventByName(app))
	}

	private := r.Group("user")
	private.Use(middleware.AuthUser())
	{
		private.GET("/rsvplist", getRSVPListForUser())
	}

	return r
}
