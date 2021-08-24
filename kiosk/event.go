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
	} else {
		c.JSON(http.StatusOK, e)
	}
}
