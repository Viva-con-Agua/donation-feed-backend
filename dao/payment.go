package dao

import "github.com/Viva-con-Agua/vcago"

type Payment struct {
	Money   vcago.Money `json:"money"`
	Contact NatsContact `json:"contact" bson:"contact"`
}
