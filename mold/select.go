package mold

import (
	_event "github.com/streambinder/peephole/mold/event"
	_util "github.com/streambinder/peephole/util"
	"github.com/xujiajun/nutsdb"
)

func (db *Mold) Select(filter string, page, limit int) ([]_event.Event, error) {
	var (
		data nutsdb.Entries
		err  error
	)

	tx, err := db.Begin(false)
	if err != nil {
		return []_event.Event{}, err
	}
	defer tx.Rollback()

	if filter != "" {
		data, _, err = tx.PrefixSearchScan(bucket, []byte{}, filter, page*limit, limit)
	} else {
		data, _, err = tx.PrefixScan(bucket, []byte{}, page*limit, limit)
	}

	if err != nil &&
		(err.Error() == "prefix scans no result" || err.Error() == "prefix and search scans not found") {
		return []_event.Event{}, nil
	} else if err != nil {
		return []_event.Event{}, err
	}

	batch := []_event.Event{}
	for _, entry := range data {
		e := _event.Event{}
		if err := _util.Unmarshal(entry.Value, &e); err != nil {
			return []_event.Event{}, err
		}
		batch = append(batch, e.Outline())
	}

	return batch, nil
}
