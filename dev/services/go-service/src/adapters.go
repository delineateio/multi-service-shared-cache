package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

const lang = "go"

func healthzAdapter(c *gin.Context) {
	addHeaders(c)
	err := pingCache()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func writeAdapter(c *gin.Context) {
	addHeaders(c)
	i := 1
	go cacheWrite(int64(i))
	c.Status(http.StatusAccepted)
}

func readAdapter(c *gin.Context) {
	addHeaders(c)
	ch := make(chan readResult)
	go cacheRead(ch)
	r := <-ch
	if r.err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count": r.value,
		})
	}
}

func addHeaders(c *gin.Context) {
	c.Header("Delineateio-Lang", lang)
	c.Header("Delineateio-Hostname", getHostname())
}

func getHostname() string {

	hostname, err := os.Hostname()
	if err != nil {
		log.Err(err).Msg("get hostname failed")
	}
	return hostname
}
