package main

import (
	"donation-feed-backend/config"
	"donation-feed-backend/dao"
	"donation-feed-backend/handlers"
	"github.com/Viva-con-Agua/vcago"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

func main() {
	cfg := config.LoadFromEnv()
	e := echo.New()
	e.HTTPErrorHandler = vcago.HTTPErrorHandler
	e.Validator = vcago.JSONValidator
	e.Use(vcago.CORS.Init())
	e.Use(vcago.Logger.Init("donation-feed-backend"))

	donationEvents := make(chan dao.ServerSentEvent[dao.DonationEvent])
	defer close(donationEvents)
	testEmitEvents(donationEvents)

	api := e.Group("/api")
	api.GET("/donations", handlers.CreateHandlerForDonationFeed(donationEvents))

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(cfg.AppPort)))
}

func testEmitEvents(eventChan chan dao.ServerSentEvent[dao.DonationEvent]) {
	ticker := time.NewTicker(5000 * time.Millisecond)
	go func() {
		for {
			<-ticker.C
			event := dao.ServerSentEvent[dao.DonationEvent]{
				ID:        0,
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
	}()
}
