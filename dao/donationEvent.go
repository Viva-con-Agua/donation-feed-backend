package dao

import "github.com/Viva-con-Agua/vcago"

// ServerSentEvent are sent on an HTTP event-stream and contain the specified data
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#event_stream_format
// for details about the event stream format
type ServerSentEvent[T any] struct {
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

	// The amount of money that was donated in this donation
	DonatedMoney vcago.Money `json:"money"`

	// The total amount of money that has been accumulated through donations in all currencies.
	//
	// The currency (e.g. "€") serves as the map key while the total amount of donated money in that currency is the
	// maps value.
	TotalMoney map[string]int64 `json:"totalMoney"`
}
