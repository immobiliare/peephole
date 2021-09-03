package kiosk

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/sirupsen/logrus"
	_mold "github.com/streambinder/peephole/mold"
	_event "github.com/streambinder/peephole/mold/event"
	_util "github.com/streambinder/peephole/util"
)

type Kiosk struct {
	mold      *_mold.Mold
	router    *gin.Engine
	eventChan chan *_event.Event
	config    *Config
	boxes     map[string]packr.Box
	minifier  minifyFS
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
	k.boxes = map[string]packr.Box{
		"static":    packr.NewBox("assets/static"),
		"templates": packr.NewBox("assets/templates"),
	}
	logrus.Println(k.boxes["static"].Path)
	k.minifier = newMinifyFS(k.boxes["static"], "/assets")

	k.router = gin.Default()
	k.router.Use(cacheControl)
	k.router.Use(gzip.Gzip(gzip.DefaultCompression))

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
