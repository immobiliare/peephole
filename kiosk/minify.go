package kiosk

import (
	"bytes"
	"io"
	"net/http"
	"path/filepath"

	"github.com/gobuffalo/packd"
	"github.com/sirupsen/logrus"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

var (
	cache     = map[string]http.File{}
	mimetypes = map[string]string{
		".html": "text/html",
		".css":  "text/css",
		".js":   "text/javascript",
	}
	minifiers = map[string]func(m *minify.M, w io.Writer, r io.Reader, params map[string]string) error{
		"text/html":              html.Minify,
		"text/css":               css.Minify,
		"text/javascript":        js.Minify,
		"application/javascript": js.Minify,
	}
)

type minifyFS struct {
	http.FileSystem
	proxy    http.FileSystem
	minifier *minify.M
}

func minifier() *minify.M {
	m := minify.New()
	for k, v := range minifiers {
		m.AddFunc(k, v)
	}
	return m
}

func newMinifyFS(fs http.FileSystem, prefix string) minifyFS {
	return minifyFS{
		proxy:    fs,
		minifier: minifier(),
	}
}

func (fs minifyFS) Open(name string) (http.File, error) {
	if f, ok := cache[name]; ok {
		logrus.WithField("name", name).Debugln("Reusing cached minified file")
		return f, nil
	}

	f, err := fs.proxy.Open(name)
	if err != nil {
		return nil, err
	}

	mimetype, ok := mimetypes[filepath.Ext(name)]
	if !ok {
		logrus.WithField("name", name).Debugln("No minifier available")
		return f, nil
	}

	logrus.WithField("name", name).Println("Minifying asset")
	minFile, err := fs.minify(filepath.Base(name), mimetype, f)
	if err != nil {
		return f, nil
	}

	cache[name] = minFile
	logrus.WithField("count", len(cache)).Println("Minified asset added to the cache")
	return minFile, nil
}

func (fs minifyFS) minify(name, mediatype string, f http.File) (http.File, error) {
	buf := &bytes.Buffer{}
	err := fs.minifier.Minify(mediatype, buf, f)
	if err != nil {
		return f, err
	}

	return packd.NewFile(name, buf)
}
