package mold

import (
	"sort"
	"sync"

	"github.com/sirupsen/logrus"
	_util "github.com/streambinder/peephole/util"
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
	config        *Config
	opGetMutex    sync.Mutex
	opGet         chan *Event
	opSelectMutex sync.Mutex
	opSelect      chan []Event
	opCountMutex  sync.Mutex
	opCount       chan int
}

func Init(config *Config) (*Mold, error) {
	opt := nutsdb.DefaultOptions
	opt.Dir = config.Spool
	db, err := nutsdb.Open(opt)
	if err != nil {
		return nil, err
	}

	mold := &Mold{
		db,
		config,
		sync.Mutex{},
		make(chan *Event),
		sync.Mutex{},
		make(chan []Event),
		sync.Mutex{},
		make(chan int),
	}
	go func() {
		logrus.WithField("count", mold.Count()).Println("DB succesfully initialized")
	}()

	return mold, nil
}

func (db *Mold) Write(e *Event) error {
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

func (db *Mold) Read(jid string) (*Event, error) {
	db.opGetMutex.Lock()
	defer db.opGetMutex.Unlock()

	go func() {
		if err := db.View(
			func(tx *nutsdb.Tx) error {
				data, err := tx.Get(bucket, []byte(jid))
				if err != nil {
					db.opGet <- nil
					return err
				}

				e := new(Event)
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

func (db *Mold) Select(filter string, page, limit int) ([]Event, error) {
	db.opSelectMutex.Lock()
	defer db.opSelectMutex.Unlock()

	go func() {
		if err := db.View(
			func(tx *nutsdb.Tx) error {
				data, err := tx.GetAll(bucket)
				if err != nil {
					db.opSelect <- []Event{}
					return err
				}

				batch := []Event{}
				for _, entry := range data {
					e := Event{}
					if err := _util.Unmarshal(entry.Value, &e); err != nil {
						db.opSelect <- []Event{}
						return err
					}

					if filter == "" || e.Match(filter) {
						batch = append(batch, e.Outline())
					}
				}

				db.opSelect <- batch
				return nil
			}); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"filter": filter,
				"limit":  limit,
			}).Errorln("Unable to select keys")
		}
	}()

	batch := <-db.opSelect
	if len(batch) < limit*page {
		return []Event{}, nil
	}

	sort.Slice(batch, func(i, j int) bool {
		return batch[i].Timestamp.Before(batch[j].Timestamp)
	})

	batch = batch[limit*page:]
	if len(batch) > limit {
		batch = batch[:limit]
	}

	return batch, nil
}

func (db *Mold) Count() int {
	db.opCountMutex.Lock()
	defer db.opCountMutex.Unlock()

	go func() {
		if err := db.View(
			func(tx *nutsdb.Tx) error {
				data, err := tx.GetAll(bucket)
				if err != nil {
					db.opCount <- 0
					return err
				}

				db.opCount <- len(data)
				return nil
			}); err != nil {
			logrus.WithError(err).Errorln("Unable to count keys")
		}
	}()

	return <-db.opCount
}
