package mold

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	_util "github.com/streambinder/peephole/util"
)

var unitToTick = map[rune]time.Duration{
	'y': 24 * time.Hour,
	'M': 24 * time.Hour,
	'd': 30 * time.Minute,
	// 'h': 5 * time.Minute,
	'h': 5 * time.Second,
	'm': 30 * time.Second,
	's': time.Second,
}

func init() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func (db *Mold) housekeep() error {
	for range time.NewTicker(unitToTick[_util.Unit(db.config.Retention)]).C {
		if err := db.expireOlder(_util.ToNow(db.config.Retention)); err != nil {
			return err
		}
		logrus.WithFields(logrus.Fields{
			"next": time.Now().Add(unitToTick[_util.Unit(db.config.Retention)]).Format(time.RFC822),
		}).Println("Planning next housekeeping")
	}
	return nil
}

func (db *Mold) expireOlder(t time.Time) error {
	var (
		iter = db.NewIter(nil)
		jids = [][]byte{}
	)

	for iter.First(); iter.Valid(); iter.Next() {
		e := new(Event)
		if err := _util.Unmarshal(iter.Value(), &e); err != nil {
			return err
		}

		if e.Timestamp.Before(t) {
			jids = append(jids, iter.Key())
			logrus.WithFields(logrus.Fields{
				"jid":       e.Jid,
				"timestamp": e.Timestamp.Format(time.RFC822),
			}).Debugln("Record set for expiration")
		}
	}

	if err := iter.Close(); err != nil {
		return err
	}

	b := db.NewBatch()
	for _, jid := range jids {
		if err := b.Delete(jid, nil); err != nil {
			logrus.WithField("jid", string(jid)).Errorln("Unable to delete record")
		} else {
			logrus.WithFields(logrus.Fields{
				"jid": string(jid),
			}).Println("Record expired")
		}
	}

	if err := b.Commit(nil); err != nil {
		return err
	}

	go func() {
		logrus.WithField("count", db.count()).Debugln("Housekeeping done")
	}()

	return nil
}
