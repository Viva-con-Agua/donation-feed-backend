package main

import (
	"donation-feed-backend/args"
	"donation-feed-backend/dao"
	"donation-feed-backend/handlers"
	"github.com/Viva-con-Agua/vcago"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

func main() {
	programArgs := args.ParseProgramArgs()
	e := echo.New()
	e.HTTPErrorHandler = vcago.HTTPErrorHandler
	e.Validator = vcago.JSONValidator
	e.Use(vcago.CORS.Init())
	e.Use(vcago.Logger.Init("donation-feed-backend"))

	donationEvents := make(chan dao.ServerSentEvent[dao.DonationEvent])
	defer close(donationEvents)

	if *programArgs.StartDummyEmitter {
		e.Logger.Info("Starting dummy emitter to emit fake donation events")
		go runDummyEmitter(donationEvents)
	}

	api := e.Group("/api")
	api.GET("/donations", handlers.CreateHandlerForDonationFeed(donationEvents))

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*programArgs.Port)))
}

// Run an emitter that periodically publishes donation events on the given channel
func runDummyEmitter(eventChan chan dao.ServerSentEvent[dao.DonationEvent]) {
	ticker := time.NewTicker(5000 * time.Millisecond)
	for {
		<-ticker.C
		event := dao.ServerSentEvent[dao.DonationEvent]{
			EventType: "donation",
			Data: dao.DonationEvent{
				Name: "finn",
				Money: vcago.Money{
					Amount:   10,
					Currency: "€",
				},
			},
		}

		// send event if possible, otherwise ignore it
		select {
		case eventChan <- event:
		default:
		}
	}
}
