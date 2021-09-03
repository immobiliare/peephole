package kiosk

import (
	"regexp"

	"github.com/gin-gonic/gin"
)

var r = regexp.MustCompile(".(js|css|ico|png)$")

func cacheControl(c *gin.Context) {
	if r.MatchString(c.Request.URL.Path) {
		c.Header("cache-control", "max-age=315360000; public")
	}
	c.Next()
}
