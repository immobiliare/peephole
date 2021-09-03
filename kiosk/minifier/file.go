package minifier

import (
	"io"
	"net/http"
)

type minFile struct {
	io.Reader
	http.File
}

func (f minFile) Close() error {
	return nil
}
