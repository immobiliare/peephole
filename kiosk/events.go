package kiosk

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_mold "github.com/streambinder/peephole/mold"
)

func (k *Kiosk) eventsHandler(c *gin.Context) {
	var (
		filter = c.Query("q")
		pPage  = c.DefaultQuery("p", "1")
		pLimit = c.DefaultQuery("l", "10")
	)

	page, err := strconv.Atoi(pPage)
	if err != nil {
		logrus.WithError(err).WithField("page", pPage).Errorln("Unable to parse page number")
	}

	limit, err := strconv.Atoi(pLimit)
	if err != nil {
		logrus.WithError(err).WithField("limit", pLimit).Errorln("Unable to parse page limit number")
	}

	if e, err := k.mold.Select(filter, page, limit); err != nil {
		logrus.WithError(err).Warnln("Unable to select events")
		c.JSON(http.StatusInternalServerError, []_mold.Event{})
	} else {
		c.JSON(http.StatusOK, e)
	}
}
