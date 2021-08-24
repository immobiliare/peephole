package kiosk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (k *Kiosk) pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
