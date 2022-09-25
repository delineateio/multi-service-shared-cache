package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type logEntry struct {
	start  time.Time
	path   string
	raw    string
	method string
	ctx    *gin.Context
}

func newLogEntry(c *gin.Context) logEntry {
	return logEntry{
		start:  time.Now(),
		path:   c.Request.URL.Path,
		raw:    c.Request.URL.RawQuery,
		method: c.Request.Method,
		ctx:    c,
	}
}

func (e *logEntry) getLatency() string {
	latency := time.Since(e.start)
	if latency > time.Minute {
		latency = latency.Truncate(time.Second)
	}
	return latency.String()
}

func (e *logEntry) getFullPath() string {
	if e.raw != "" {
		return fmt.Sprintf("%s?%s", e.path, e.raw)
	} else {
		return e.path
	}
}

func (e *logEntry) next() {
	e.ctx.Next()
}

func (e *logEntry) getStatus() int {
	return e.ctx.Writer.Status()
}

func (e *logEntry) getMessage() string {
	return e.ctx.Errors.ByType(gin.ErrorTypePrivate).String()
}

func (e *logEntry) write(logger *zerolog.Logger) {
	var event *zerolog.Event
	if e.getStatus() >= 500 {
		event = logger.Error()
	} else {
		event = logger.Info()
	}
	event.Str("method", e.method).
		Int("status_code", e.getStatus()).
		Str("path", e.getFullPath()).
		Str("latency", e.getLatency()).
		Msg(e.getMessage())
}
