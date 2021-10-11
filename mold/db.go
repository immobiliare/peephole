package mold

import (
	"github.com/sirupsen/logrus"
	"github.com/xujiajun/nutsdb"
)

/*
 * module requirements include:
 * - persistence
 * - auto-expiration
 */

const (
	bucket = "_events"
)

type Mold struct {
	*nutsdb.DB
	config *Config
}

func Init(config *Config) (*Mold, error) {
	opt := nutsdb.DefaultOptions
	opt.Dir = config.Spool
	db, err := nutsdb.Open(opt)
	if err != nil {
		return nil, err
	}

	mold := &Mold{db, config}
	go func() {
		if count, err := mold.Count(""); err != nil {
			logrus.WithError(err).Errorln("Unable to cound DB entries")
		} else {
			logrus.WithField("count", count).Println("DB succesfully initialized")
		}
	}()

	return mold, nil
}
