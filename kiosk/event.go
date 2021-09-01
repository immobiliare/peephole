package kiosk

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (k *Kiosk) eventHandler(c *gin.Context) {
	if e, err := k.mold.Read(c.Param("jid")); err != nil {
		logrus.WithError(err).WithField("jid", c.Param("jid")).Errorln("Unable to query event")
		c.Status(http.StatusNotFound)
	} else if e == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
		c.String(http.StatusOK, e.Raw)
	}
}
