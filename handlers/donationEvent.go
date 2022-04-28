package handlers

import (
	"donation-feed-backend/dao"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateHandlerForDonationFeed(eventChan chan dao.ServerSentEvent[dao.DonationEvent]) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		responseWriter := c.Response().Writer
		responseWriter.Header().Set("Content-Type", "text/event-stream")
		responseWriter.Header().Set("Cache-Control", "no-store")
		responseWriter.Header().Set("Connection", "keep-alive")
		responseWriter.WriteHeader(http.StatusOK)

		for {
			select {
			case event := <-eventChan:
				if err := writeEvent(responseWriter, &event); err != nil {
					return err
				}
				c.Response().Flush()

			case <-c.Request().Context().Done():
				return nil
			}
		}
	}
}

func writeEvent(writer http.ResponseWriter, event *dao.ServerSentEvent[dao.DonationEvent]) (err error) {
	// write "event" field
	if _, err := fmt.Fprintf(writer, "event: %s\n", event.EventType); err != nil {
		return err
	}
	// write "data field"
	if _, err = fmt.Fprint(writer, "data: "); err != nil {
		return err
	}
	enc := json.NewEncoder(writer)
	if err = enc.Encode(event.Data); err != nil {
		return err
	}
	// write separator
	if _, err = fmt.Fprint(writer, "\n\n"); err != nil {
		return err
	}
	return nil
}
