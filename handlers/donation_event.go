package handlers

import (
	"donation-feed-backend/dao"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateHandlerForDonationFeed(eventChan chan dao.DonationEvent) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		responseWriter := c.Response().Writer
		responseWriter.Header().Set("Content-Type", "text/event-stream")
		responseWriter.Header().Set("Cache-Control", "no-store")
		responseWriter.Header().Set("Connection", "keep-alive")

		for {
			select {
			case event := <-eventChan:
				if err := c.Echo().JSONSerializer.Serialize(c, event, ""); err != nil {
					return err
				}
				responseWriter.(http.Flusher).Flush()

			case <-c.Request().Context().Done():
				return
			}
		}
	}
}
