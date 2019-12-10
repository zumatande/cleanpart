package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.LevelFieldName = "severity"
	err := run()
	if err != nil {
		// [Q] should we wrap errors here?
		log.Error().Err(err).Msg("error starting partner service")
		os.Exit(1)
	}
}

// we put main logic inside this function so it is treated
// like any other function in the service, testable
func run() (err error) {
	// spin up initialization and dependency injection
	return err
}
