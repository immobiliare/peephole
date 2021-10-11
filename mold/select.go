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
	if err != nil {
		return []_event.Event{}, err
	}

	batch := []_event.Event{}
	for _, entry := range data {
		e := _event.Event{}
		if err := _util.Unmarshal(entry.Value, &e); err != nil {
			return []_event.Event{}, err
		}

		if filter == "" || e.Match(filter) {
			batch = append(batch, e.Outline())
		}
	}

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
