package kiosk

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (k *Kiosk) indexHandler(c *gin.Context) {
	// TODO:
	// - use minified template
	// - render html as template
	if html, err := k.boxes["templates"].FindString("index.html"); err != nil {
		logrus.WithError(err).WithField("jid", c.Param("jid")).Errorln("Unable to query event")
		c.Status(http.StatusNotFound)
	} else {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.Next()
		c.String(http.StatusOK, html)
	}
}
