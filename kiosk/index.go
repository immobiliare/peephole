package kiosk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (k *Kiosk) indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Peephole"})
}
