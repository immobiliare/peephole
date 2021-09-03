package minifier

import (
	"bytes"
	"fmt"
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

type Minifier FS

type FS struct {
	http.FileSystem
	minifier *minify.M
	proxy    http.FileSystem
	cache    map[string][]byte
}

func Init(proxy http.FileSystem) *FS {
	fs := &FS{proxy: proxy}
	fs.cache = make(map[string][]byte)
	fs.minifier = minify.New()
	for k, v := range minifiers {
		fs.minifier.AddFunc(k, v)
	}
	return fs
}

func (fs FS) Open(name string) (http.File, error) {
	// for k, _ := range fs.cache {
	// 	logrus.Println("cache entry " + k)
	// 	info, err := v.Stat()
	// 	if err != nil {
	// 		logrus.WithField("k", k).Debugln("Unable to read cache entry")
	// 	} else {
	// 		buf := make([]byte, 30)
	// 		v.Read(buf)
	// 		logrus.WithFields(logrus.Fields{
	// 			"k":      k,
	// 			"len":    info.Size(),
	// 			"prefix": string(buf),
	// 		}).Debugln("Cache entry found")
	// 	}
	// }

	if min, ok := fs.cache[name]; ok {
		logrus.WithField("name", name).Debugln("Reusing cached minified asset")
		return packd.NewFile(name, bytes.NewReader(min))
	}

	f, err := fs.proxy.Open(name)
	if err != nil {
		return nil, err
	}

	min, err := fs.Minify(f)
	if err != nil {
		return f, err
	}

	return min, nil
}

func (fs FS) Minify(f http.File) (http.File, error) {
	info, err := f.Stat()
	if err != nil {
		return f, err
	}

	name := fmt.Sprintf("/%s", info.Name())
	mimetype, ok := mimetypes[filepath.Ext(name)]
	if !ok {
		logrus.WithField("name", name).Debugln("No minifier available")
		return f, nil
	}

	logrus.WithField("name", name).Println("Minifying asset")
	min, err := fs.minify(filepath.Base(name), mimetype, f)
	if err != nil {
		return f, nil
	}

	return min, nil
}

func (fs FS) minify(name, mediatype string, f http.File) (http.File, error) {
	buf := &bytes.Buffer{}
	err := fs.minifier.Minify(mediatype, buf, f)
	if err != nil {
		return f, err
	}

	fs.cache["/"+name] = buf.Bytes()
	return packd.NewFile(name, bytes.NewReader(buf.Bytes()))
}
