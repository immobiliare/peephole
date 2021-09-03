package kiosk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (k *Kiosk) faviconHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/assets/favicon.ico")
}
