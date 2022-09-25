package main

import (
	"github.com/rs/zerolog/log"
)

func init() {
	load()
	watchConfig()
}

func load() {
	err := loadConfig()
	if err != nil {
		log.Err(err).Msg("Failed to load config")
	}
	err = setLogLevel()
	if err != nil {
		log.Err(err).Msg("Failed to set log level")
	}
	err = connectToCache()
	if err != nil {
		log.Err(err).Msg("Failed to connect to cache")
	}
}
