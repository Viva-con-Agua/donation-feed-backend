package dao

import "github.com/Viva-con-Agua/vcago"

// ServerSentEvent are sent on an HTTP event-stream and contain the specified data
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#event_stream_format
// for details about the event stream format
type ServerSentEvent[T any] struct {
	// Monotonically incrementing number that can be used to identify events
	ID int
	// Which type of event this represents.
	// Generally this is dependent on the type of the data.
	EventType string
	// Actual event data
	Data T
}

type DonationEvent struct {
	// An optional name that can be displayed to show who donated.
	// Can also be nil
	Name string `json:"name"`

	// The amount of money that was donated
	Money vcago.Money `json:"money"`
}
