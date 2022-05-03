package nats

import (
	"github.com/Viva-con-Agua/donation-feed-backend/dao"
	"github.com/nats-io/nats.go"
)

const TopicPaymentDone = "payment.done"

func SubscribeToPayments(conn *nats.EncodedConn, handler func(payment *dao.Payment)) (*nats.Subscription, error) {
	return conn.Subscribe(TopicPaymentDone, handler)
}
