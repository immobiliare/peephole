package mold

import (
	"sort"

	_event "github.com/streambinder/peephole/mold/event"
	_util "github.com/streambinder/peephole/util"
)

func (db *Mold) Select(filter string, page, limit int) ([]_event.Event, error) {
	tx, err := db.Begin(false)
	if err != nil {
		return []_event.Event{}, err
	}
	defer tx.Rollback()

	data, err := tx.GetAll(bucket)
	// data, _, err := tx.PrefixSearchScan(bucket, []byte(""), "*", 0, 100)
	if err != nil {
		return []_event.Event{}, err
	}

	batch := []_event.Event{}
	for i := len(data) - 1; i >= 0; i-- {
		entry := data[i]
		e := _event.Event{}
		if err := _util.Unmarshal(entry.Value, &e); err != nil {
			return []_event.Event{}, err
		}

		if filter == "" || e.Match(filter) {
			batch = append(batch, e.Outline())
			if len(batch) > limit*page+limit {
				break
			}
		}
	}

	if len(batch) > limit*page {
		batch = batch[limit*page:]
	}

	if len(batch) > limit {
		batch = batch[:limit]
	}

	sort.Slice(batch, func(i, j int) bool {
		return batch[i].Timestamp.Before(batch[j].Timestamp)
	})

	return batch, nil
}
