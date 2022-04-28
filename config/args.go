package config

import (
	"flag"
	"github.com/Viva-con-Agua/vcago"
)

type ProgramArguments struct {
	StartDummyEmitter *bool
}

func ParseProgramArgs() ProgramArguments {
	args := ProgramArguments{}

	args.StartDummyEmitter = flag.Bool(
		"start-dummy-emitter",
		vcago.Config.GetEnvBool("APP_START_DUMMY_EMITTER", "n", false),
		"Whether a dummy emitter should be started that emits a dummy donation events automatically",
	)

	flag.Parse()

	return args
}
