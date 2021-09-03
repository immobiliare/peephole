package kiosk

import (
	"regexp"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/sirupsen/logrus"
	_min "github.com/streambinder/peephole/kiosk/minifier"
	_mold "github.com/streambinder/peephole/mold"
	_event "github.com/streambinder/peephole/mold/event"
	_util "github.com/streambinder/peephole/util"
)

type Kiosk struct {
	mold      *_mold.Mold
	router    *gin.Engine
	eventChan chan *_event.Event
	config    *Config
	templates packr.Box
	minifier  *_min.FS
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
	k.templates = packr.NewBox("assets/templates")
	k.minifier = _min.Init(packr.NewBox("assets/static"))

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
	group.StaticFS("/assets", k.minifier)
	group.GET("/", k.indexHandler)
	group.GET("/events", k.eventsHandler)
	group.GET("/events/:jid", k.eventHandler)

	return k
}

func (k *Kiosk) Serve() error {
	return k.router.Run(k.config.Bind)
}
