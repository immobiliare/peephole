package kiosk

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (k *Kiosk) streamHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Next()
	c.Stream(func(w io.Writer) bool {
		if e, ok := <-k.eventChan; ok {
			msg := e.Outline()
			logrus.WithField("jid", msg.Jid).Debugln("Sending SSE message")
			c.SSEvent("event", msg)
			return true
		}
		return false
	})
}
