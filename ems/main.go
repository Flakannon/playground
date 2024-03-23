package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/playground/ems/database/eventdb"
	"github.com/playground/ems/restapi"
)

func main() {
	var app restapi.App

	app.EventFinder = eventdb.EventDBMock{}
	engine := restapi.NewEngine(app)

	server := http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error with listen and serve: ", err.Error())
		}
	}()

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Error shutting down gracefully", err.Error())
	}

	log.Print("Server exited safely")
}

// event managament system
// Event Creation: Users must be able to create events by providing details such as name, description, location, start date/time, and end date/time.
// post endpoint for create event

// Event Modification: Users should be able to update the details of an event.
// patch to modify event

// Event Deletion: Users must have the ability to delete an event.
// delete to remove event

// Viewing Events: There should be functionality to view all events and view details of a specific event.
// get for events can get all events, by id, by time for filter?

// RSVP to Events: Users can RSVP to events, indicating their intention to attend. The system should track these RSVPs.
// Patch to rsvp for specific user

// User Authentication: Implement a simple user authentication mechanism to identify users.

// Permissions: Only the event creator can modify or delete the event. However, any authenticated user can view events and RSVP.
