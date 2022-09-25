package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type server struct {
	router *gin.Engine
}

func newServer() *server {

	gin.SetMode(getString("web.mode", "release"))
	router := gin.New()
	router.Use(newStructuredLogger())
	router.Use(gin.Recovery())
	return &server{
		router: router,
	}
}

func (s *server) addGet(path string, f func(*gin.Context)) {
	s.router.GET(path, f)
}

func (s *server) addPost(path string, f func(*gin.Context)) {
	s.router.POST(path, f)
}

func (s *server) start() {
	port := ":" + getString("web.port", "8080")
	log.Info().Msgf("server listening on %s", port)
	s.router.Run(port)
}
