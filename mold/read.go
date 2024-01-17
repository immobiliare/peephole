package mold

import (
	_event "github.com/immobiliare/peephole/mold/event"
	_util "github.com/immobiliare/peephole/util"
	"github.com/sirupsen/logrus"
)

func (db *Mold) Read(id string) (*_event.Event, error) {
	tx, err := db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := tx.Rollback()
		if err != nil {
			logrus.Debugln(err.Error())
		}
	}()

	data, err := tx.Get(bucket, []byte(id))
	if err != nil && err.Error() == "key not found" {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	e := new(_event.Event)
	if err := _util.Unmarshal(data.Value, e); err != nil {
		return nil, err
	}

	return e, nil
}
