package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func getLogLevel() string {
	return getString("log.level", "info")
}

func setLogLevel() error {

	level := getLogLevel()
	l, err := zerolog.ParseLevel(level)
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(l)
	log.Info().Msgf("set of log level successful: %s", level)
	return nil
}
