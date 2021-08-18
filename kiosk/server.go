package kiosk

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_mold "gitlab.rete.farm/dpucci/peephole/mold"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Serve() error {
	r := gin.Default()

	r.LoadHTMLGlob("kiosk/assets/templates/*html")
	r.Static("/assets", "kiosk/assets/static")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Peephole",
		})
	})
	r.GET("/events", func(c *gin.Context) {
		if e, err := _mold.Select(); err != nil {
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, e)
		}
	})

	go func(r *gin.Engine) {
		if err := r.Run(); err != nil {
			logrus.WithError(err).Fatalln("Unable to start kiosk server")
		}
	}(r)

	return nil
}
