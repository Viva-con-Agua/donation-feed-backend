package db

import (
	"github.com/Viva-con-Agua/vcago"
)

func SetupDb() (db *vcago.MongoDB, coll *vcago.MongoColl) {
	db = vcago.NewMongoDB("donation-feed-backend")
	coll = db.Collection("donations")
	return
}
