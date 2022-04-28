package dao

import "github.com/Viva-con-Agua/vcago"

// DonationEvent is the type that is emitted by the HTTP event stream
type DonationEvent struct {
	// An optional name that can be displayed to show who donated.
	// Can also be nil
	Name string `json:"name"`

	// The amount of money that was donated
	Money vcago.Money `json:"money"`
}
