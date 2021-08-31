package mold

import (
	"sort"

	"github.com/sirupsen/logrus"
	_event "github.com/streambinder/peephole/mold/event"
	_util "github.com/streambinder/peephole/util"
	"github.com/xujiajun/nutsdb"
)

func (db *Mold) Select(filter string, page, limit int) ([]_event.Event, error) {
	db.opSelectMutex.Lock()
	defer db.opSelectMutex.Unlock()

	go func() {
		if err := db.View(
			func(tx *nutsdb.Tx) error {
				data, err := tx.GetAll(bucket)
				if err != nil {
					db.opSelect <- []_event.Event{}
					return err
				}

				batch := []_event.Event{}
				for _, entry := range data {
					e := _event.Event{}
					if err := _util.Unmarshal(entry.Value, &e); err != nil {
						db.opSelect <- []_event.Event{}
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
		return []_event.Event{}, nil
	}

	sort.Slice(batch, func(i, j int) bool {
		return batch[i].Timestamp.After(batch[j].Timestamp)
	})

	batch = batch[limit*page:]
	if len(batch) > limit {
		batch = batch[:limit]
	}

	return batch, nil
}
