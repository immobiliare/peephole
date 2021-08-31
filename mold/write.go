package mold

import (
	_event "github.com/streambinder/peephole/mold/event"
	_util "github.com/streambinder/peephole/util"
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
				return tx.Put(bucket, []byte(e.Jid), bytes, r)
			}
		})
}
