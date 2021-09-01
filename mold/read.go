package mold

import (
	"github.com/sirupsen/logrus"
	_event "github.com/streambinder/peephole/mold/event"
	_util "github.com/streambinder/peephole/util"
	"github.com/xujiajun/nutsdb"
)

func (db *Mold) Read(jid string) (*_event.Event, error) {
	db.opGetMutex.Lock()
	defer db.opGetMutex.Unlock()

	go func() {
		if err := db.View(
			func(tx *nutsdb.Tx) error {
				data, err := tx.Get(bucket, []byte(jid))
				if err != nil && err.Error() == "key not found" {
					db.opGet <- nil
					return nil
				} else if err != nil {
					db.opGet <- nil
					return err
				}

				e := new(_event.Event)
				if err := _util.Unmarshal(data.Value, e); err != nil {
					db.opGet <- nil
					return err
				}

				db.opGet <- e
				return nil
			}); err != nil {
			logrus.WithError(err).WithField("jid", jid).Errorln("Unable to read key")
		}
	}()

	return <-db.opGet, nil
}
