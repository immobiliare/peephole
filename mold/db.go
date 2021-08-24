package mold

import (
	"sort"
	"sync"

	"github.com/cockroachdb/pebble"
	"github.com/sirupsen/logrus"
	_config "github.com/streambinder/peephole/config"
	_util "github.com/streambinder/peephole/util"
)

type Mold struct {
	*pebble.DB
	config *_config.Mold
	mutex  sync.Mutex
}

func Init(config *_config.Mold) (*Mold, error) {
	db, err := pebble.Open(config.Spool, &pebble.Options{})
	if err != nil {
		return nil, err
	}

	mold := &Mold{db, config, sync.Mutex{}}
	go func() {
		if mold.housekeep(); err != nil {
			logrus.WithError(err).Errorln("Unable to enforce db retention")
		}
	}()
	go func() {
		logrus.WithField("count", mold.count()).Println("DB succesfully initialized")
	}()
	return mold, nil
}

func (db *Mold) Write(e *Event) error {
	bytes, err := _util.Marshal(e)
	if err != nil {
		return err
	}

	return db.Set([]byte(e.Jid), bytes, nil)
}

func (db *Mold) Read(jid string) (*Event, error) {
	value, closer, err := db.Get([]byte(jid))
	if err != nil {
		return nil, err
	}

	if err := closer.Close(); err != nil {
		return nil, err
	}

	e := new(Event)
	if err := _util.Unmarshal(value, e); err != nil {
		return nil, err
	}

	return e, nil
}

func (db *Mold) Select(filter string, limit int) ([]Event, error) {
	var (
		iter  = db.NewIter(nil)
		batch = []Event{}
	)

	for iter.First(); iter.Valid(); iter.Next() {
		e := Event{}
		if err := _util.Unmarshal(iter.Value(), &e); err != nil {
			return []Event{}, err
		}

		if filter == "" || e.Match(filter) {
			batch = append(batch, e)
		}
	}

	if err := iter.Close(); err != nil {
		return []Event{}, err
	}

	sort.Slice(batch, func(i, j int) bool {
		return batch[i].Timestamp.Before(batch[j].Timestamp)
	})

	if len(batch) > limit {
		batch = batch[:limit]
	}

	return batch, nil
}

func (db *Mold) count() int {
	var (
		iter  = db.NewIter(nil)
		count = 0
	)
	defer iter.Close()

	for iter.First(); iter.Valid(); iter.Next() {
		count++
	}

	return count
}
