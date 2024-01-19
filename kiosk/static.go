package kiosk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (k *Kiosk) staticHandler(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		filename = "assets/index.html"
	} else {
		filename = "assets/" + filename
	}

	blob, err := assets.ReadFile(filename)
	if err != nil {
		c.Status(http.StatusNotFound)
	}

	contentType := http.DetectContentType(blob)
	if filename == "assets/style.css" {
		contentType = "text/css; charset=utf-8"
	}

	c.Data(http.StatusOK, contentType, blob)
}
