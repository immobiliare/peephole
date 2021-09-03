package kiosk

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (k *Kiosk) indexHandler(c *gin.Context) {
	// TODO:
	// - render html as template
	const (
		template = "index.html"
		mimetype = "text/html"
	)

	html, err := k.templates.Find(template)
	if err != nil {
		logrus.WithError(err).WithField("jid", c.Param("jid")).Errorln("Unable to query event")
		c.Status(http.StatusNotFound)
	}

	min, err := k.minifier.Minify(template, html)
	if err != nil {
		logrus.WithError(err).WithField("jid", c.Param("jid")).Errorln("Unable to minify template")
		c.Status(http.StatusNotFound)
	}

	c.Data(http.StatusOK, mimetype, min)
}
