package nats

import (
	"github.com/Viva-con-Agua/donation-feed-backend/args"
	"github.com/labstack/gommon/log"
	"github.com/nats-io/nats.go"
)

func Connect(programArgs *args.ProgramArguments) (encodedConn *nats.EncodedConn, err error) {
	log.Infof("Connecting to NATS %s", *programArgs.NatsUrl)
	var conn *nats.Conn
	if conn, err = nats.Connect(*programArgs.NatsUrl); err != nil {
		return nil, err
	}

	if encodedConn, err = nats.NewEncodedConn(conn, nats.JSON_ENCODER); err != nil {
		return nil, err
	}

	log.Info("Successfully connected to NATS")
	return
}
