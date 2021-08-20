package kiosk

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_config "gitlab.rete.farm/dpucci/peephole/config"
	_mold "gitlab.rete.farm/dpucci/peephole/mold"
)

type Kiosk struct {
	router    *gin.Engine
	eventChan chan *_mold.Event
	config    *_config.Kiosk
}

func init() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func Init(eventChan chan *_mold.Event, config *_config.Kiosk) *Kiosk {
	k := new(Kiosk)
	k.config = config
	k.router = gin.Default()
	k.eventChan = eventChan
	k.router.LoadHTMLGlob("kiosk/assets/templates/*html")
	k.router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	_priv := k.router.Group("/", gin.BasicAuth(gin.Accounts(config.BasicAuth)))
	_priv.Static("/assets", "kiosk/assets/static")
	_priv.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", gin.H{"title": "Peephole"}) })
	_priv.GET("/events", func(c *gin.Context) {
		if e, err := _mold.Select(15); err != nil {
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, e)
		}
	})
	k.router.GET("/stream", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
		c.Stream(func(w io.Writer) bool {
			if e, ok := <-k.eventChan; ok {
				msg := e.Outline()
				logrus.WithField("jid", msg.Jid).Debugln("Sending SSE message")
				c.SSEvent("event", msg)
				return true
			}
			return false
		})
	})

	return k
}

func (k *Kiosk) Serve() error {
	return k.router.Run(k.config.Bind)
}
