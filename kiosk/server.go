package kiosk

import (
	"github.com/gin-gonic/gin"
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
	k.router = gin.Default()
	k.eventChan = eventChan
	k.router.LoadHTMLGlob("kiosk/assets/templates/*html")
	k.router.GET("/ping", k.pingHandler)
	k.router.GET("/stream", k.streamHandler)

	_priv := k.router.Group("/", gin.BasicAuth(gin.Accounts(config.BasicAuth)))
	_priv.Static("/assets", "kiosk/assets/static")
	_priv.GET("/", k.indexHandler)
	_priv.GET("/events", k.eventsHandler)
	_priv.GET("/events/:jid", k.eventHandler)

	return k
}

func (k *Kiosk) Serve() error {
	return k.router.Run(k.config.Bind)
}
