package mold

import (
	_event "github.com/immobiliare/peephole/mold/event"
	_util "github.com/immobiliare/peephole/util"
	"github.com/xujiajun/nutsdb"
)

func (db *Mold) Write(e *_event.Event) error {
	bytes, err := _util.Marshal(e)
	if err != nil {
		return err
	}

	return db.Update(
		func(tx *nutsdb.Tx) error {
			r, err := _util.RetentionSeconds(db.config.Retention)
			if err != nil {
				return err
			} else {
				return tx.PutWithTimestamp(bucket, []byte(e.ID), bytes, r, uint64(e.Timestamp.Unix()))
			}
		})
}
