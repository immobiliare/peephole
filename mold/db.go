package mold

import (
	"fmt"
	"sort"
)

var (
	history []*Event = []*Event{}
)

func Persist(e *Event) error {
	history = append(history, e)
	return nil
}

func Get(jid string) (*Event, error) {
	for _, e := range history {
		if e.Jid == jid {
			return e, nil
		}
	}
	return nil, fmt.Errorf("event with jid %s not found", jid)
}

func Select(limit int) ([]Event, error) {
	chunk := []Event{}
	for i := 0; i < limit; i++ {
		if i == len(history) {
			break
		}
		chunk = append(chunk, history[i].Outline())
	}
	sort.Slice(chunk, func(i, j int) bool {
		return chunk[i].Timestamp.Before(chunk[j].Timestamp)
	})
	return chunk, nil
}
