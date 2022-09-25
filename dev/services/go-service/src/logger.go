package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func newStructuredLogger() gin.HandlerFunc {
	return structuredLogger(&log.Logger)
}

func structuredLogger(logger *zerolog.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		e := newLogEntry(c)
		e.next()
		e.write(logger)
	}
}
