package main

import (
	"context"
	"fmt"
	"github.com/Viva-con-Agua/donation-feed-backend/args"
	"github.com/Viva-con-Agua/donation-feed-backend/broadcastChannel"
	"github.com/Viva-con-Agua/donation-feed-backend/dao"
	"github.com/Viva-con-Agua/donation-feed-backend/db"
	"github.com/Viva-con-Agua/donation-feed-backend/handlers"
	"github.com/Viva-con-Agua/donation-feed-backend/nats"
	"github.com/Viva-con-Agua/vcago"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	upstreamNats "github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
	"time"
	"math/rand"
)

func main() {
	programArgs := args.ParseProgramArgs()
	e := echo.New()
	e.HTTPErrorHandler = vcago.HTTPErrorHandler
	e.Validator = vcago.JSONValidator
	e.Use(vcago.CORS.Init())
	e.Use(vcago.Logger.Init("donation-feed-backend"))
	// setup required things
	eventSourceChan := make(chan dao.ServerSentEvent[dao.DonationEvent])
	eventBroadcastChan := broadcastChannel.NewBroadcastChannel(eventSourceChan)
	_, dbCollection := db.SetupDb()

	var natsConn *upstreamNats.EncodedConn
	if *programArgs.StartDummyEmitter {
		log.Info("Starting dummy emitter to emit fake donation events instead of NATS subscriber")
		go runDummyEmitter(createPaymentEventHandler(eventSourceChan, dbCollection))
	} else {
		natsConn, _ = nats.Connect(&programArgs)
		_, _ = nats.SubscribeToPayments(natsConn, createPaymentEventHandler(eventSourceChan, dbCollection))
	}

	// setup http routes
	e.GET("/v1/donation-events", handlers.CreateHandlerForDonationFeed(eventBroadcastChan))

	// start server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*programArgs.Port)))
}

// Run an emitter that periodically publishes donation events on the given channel
func runDummyEmitter(natsPaymentHandler func(payment *dao.Payment)) {
	ticker := time.NewTicker(3000 * time.Millisecond)
	totalMoney := make(map[string]int64)
	totalMoney["€"] = 0

	for {
		<-ticker.C
		natsPaymentHandler(&dao.Payment{
			Money: vcago.Money{
				Amount:   int64(rand.Intn((10000-500) + 500)),
				Currency: "€",
			},
			Contact: dao.Contact{
				FirstName: randSeq(6),
				LastName:  randSeq(6),
			},
		})
	}
}
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


// Create a handler that is able to handle new payments from NATS and process them into this application
func createPaymentEventHandler(eventChan chan dao.ServerSentEvent[dao.DonationEvent], db *vcago.MongoColl) func(payment *dao.Payment) {
	return func(payment *dao.Payment) {
		if err := db.InsertOne(context.Background(), payment); err != nil {
			log.Error(err.Error())
			return
		}

		var totalMoney []dao.AggregatedTotalMoney
		aggregateQuery := mongo.Pipeline{
			{
				{"$group", bson.D{
					{"_id", "$money.currency"},
					{"totalDonationAmount", bson.D{
						{"$sum", "$money.amount"},
					}},
				}},
			},
		}
		if err := db.Aggregate(context.Background(), aggregateQuery, &totalMoney); err != nil {
			log.Error(err.Error())
			return
		}

		// send event if possible, otherwise ignore it
		event := dao.ServerSentEvent[dao.DonationEvent]{
			EventType: "donation",
			Data: dao.DonationEvent{
				Name:         strings.TrimSpace(fmt.Sprintf("%s %s", payment.Contact.FirstName, payment.Contact.LastName)),
				DonatedMoney: payment.Money,
				TotalMoney:   dao.AggregatedTotalMoneyToMap(totalMoney),
			},
		}
		select {
		case eventChan <- event:
		default:
		}
	}
}
