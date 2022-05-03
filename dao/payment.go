package dao

import "github.com/Viva-con-Agua/vcago"

// Payment is a subset of payment information as it is published on NATS
type Payment struct {
	Money   vcago.Money `json:"money" bson:"money"`
	Contact Contact     `json:"contact" bson:"contact"`
}
