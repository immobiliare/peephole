package util

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io"
	"log"
)

// Marshal serializes into a GZIP compressed byte array
func Marshal(obj interface{}) ([]byte, error) {
	plainBuf := bytes.Buffer{}
	plainEnc := gob.NewEncoder(&plainBuf)
	if err := plainEnc.Encode(obj); err != nil {
		return []byte{}, err
	}

	gzipBuf := bytes.Buffer{}
	gzipWriter := gzip.NewWriter(&gzipBuf)
	if _, err := gzipWriter.Write(plainBuf.Bytes()); err != nil {
		return []byte{}, err
	}

	if err := gzipWriter.Close(); err != nil {
		return []byte{}, err
	}

	return gzipBuf.Bytes(), nil
}

// Unmarshal loads object from given GZIP compressed bytes array
func Unmarshal(ser []byte, object interface{}) error {
	gzipReader, err := gzip.NewReader(bytes.NewReader(ser))
	if err != nil {
		return err
	}

	data, err := io.ReadAll(io.Reader(gzipReader))
	if err != nil {
		log.Fatal(err)
	}

	if err := gzipReader.Close(); err != nil {
		return err
	}

	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(object); err != nil {
		return err
	}

	return nil
}
