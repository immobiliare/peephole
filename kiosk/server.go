package kiosk

import (
	"embed"
	"regexp"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_mold "github.com/immobiliare/peephole/mold"
	_event "github.com/immobiliare/peephole/mold/event"
	_util "github.com/immobiliare/peephole/util"
	"github.com/sirupsen/logrus"
)

//go:embed assets
var assets embed.FS

type Kiosk struct {
	mold      *_mold.Mold
	router    *gin.Engine
	eventChan chan *_event.Event
	config    *Config
}

func init() {
	if _util.Debugging() {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func Init(db *_mold.Mold, eventChan chan *_event.Event, config *Config) *Kiosk {
	k := new(Kiosk)
	k.mold = db
	k.config = config
	k.eventChan = eventChan

	k.router = gin.Default()
	k.router.Use(gzip.Gzip(gzip.DefaultCompression))
	k.router.Use(func(c *gin.Context) {
		if regexp.MustCompile(".(js|css|ico|png)$").MatchString(c.Request.URL.Path) {
			c.Header("cache-control", "max-age=315360000; public")
		}
		c.Next()
	})

	k.router.GET("/ping", k.pingHandler)
	k.router.GET("/stream", k.streamHandler)
	k.router.GET("/favicon.ico", k.faviconHandler)

	group := k.router.Group("/")
	if len(config.BasicAuth) > 0 {
		group = k.router.Group("/", gin.BasicAuth(gin.Accounts(config.BasicAuth)))
	}

	group.GET("/assets/:filename", k.staticHandler)
	group.GET("/", k.staticHandler)
	group.GET("/events", k.eventsHandler)
	group.GET("/events/:jid", k.eventHandler)

	return k
}

func (k *Kiosk) Serve() error {
	return k.router.Run(k.config.Bind)
}
