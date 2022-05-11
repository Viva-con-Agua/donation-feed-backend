package nats

import (
	"github.com/Viva-con-Agua/donation-feed-backend/args"
	"github.com/Viva-con-Agua/vcago"
	"github.com/labstack/gommon/log"
	"github.com/nats-io/nats.go"
)

var NatsHost = vcago.Config.GetEnvString("NATS_HOST", "w", "localhost")
var NatsPort = vcago.Config.GetEnvString("NATS_PORT", "w", "4222")
var NatsURL = "nats://" + NatsHost + ":" + NatsPort

func Connect(programArgs *args.ProgramArguments) (encodedConn *nats.EncodedConn, err error) {
	log.Infof("Connecting to NATS %s", *programArgs.NatsUrl)
	var conn *nats.Conn
	if conn, err = nats.Connect(NatsURL); err != nil {
		return nil, err
	}

	if encodedConn, err = nats.NewEncodedConn(conn, nats.JSON_ENCODER); err != nil {
		return nil, err
	}

	log.Info("Successfully connected to NATS")
	return
}
