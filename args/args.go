package args

import (
	"flag"
	"github.com/Viva-con-Agua/vcago"
)

type ProgramArguments struct {
	StartDummyEmitter *bool
	Port              *int
	NatsUrl           *string
}

func ParseProgramArgs() ProgramArguments {
	args := ProgramArguments{}

	args.StartDummyEmitter = flag.Bool(
		"start-dummy-emitter",
		vcago.Config.GetEnvBool("APP_START_DUMMY_EMITTER", "n", false),
		"Whether a dummy emitter should be started that emits a dummy donation events automatically",
	)
	args.Port = flag.Int(
		"port",
		vcago.Config.GetEnvInt("APP_PORT", "n", 8080),
		"On which port that application should listen",
	)
	args.NatsUrl = flag.String(
		"nats-url",
		vcago.Config.GetEnvString("APP_NATS_URL", "e", ""),
		"A nats:// url that is used to connect to a running NATS server",
	)

	flag.Parse()

	return args
}
