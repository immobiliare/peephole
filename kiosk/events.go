package kiosk

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_event "github.com/streambinder/peephole/mold/event"
)

type Pagination struct {
	Events  []_event.Event
	Page    int
	Limit   int
	HasNext bool
}

func (k *Kiosk) eventsHandler(c *gin.Context) {
	var (
		filter = c.Query("q")
		pPage  = c.DefaultQuery("p", "1")
		pLimit = c.DefaultQuery("l", "7")
	)

	p, err := strconv.Atoi(pPage)
	if err != nil {
		logrus.WithError(err).WithField("page", pPage).Errorln("Unable to parse page number")
	}

	l, err := strconv.Atoi(pLimit)
	if err != nil {
		logrus.WithError(err).WithField("limit", pLimit).Errorln("Unable to parse page limit number")
	}

	if e, err := k.mold.Select(filter, p, l); err != nil {
		logrus.WithError(err).Warnln("Unable to select events")
		c.JSON(http.StatusInternalServerError, []_event.Event{})
	} else {
		cnt, err := k.mold.Count()
		if err != nil {
			logrus.WithError(err).Errorln("Unable to count DB entries")
		}
		c.JSON(http.StatusOK, Pagination{e, p, l, len(e) > 0 && (p+1)*l < cnt})
	}
}
