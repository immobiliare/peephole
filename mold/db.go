package mold

import "sort"

var (
	history []*Event = []*Event{}
)

func Persist(e *Event) error {
	history = append(history, e)
	return nil
}

func Select(chunk int) ([]*Event, error) {
	sort.Slice(history, func(i, j int) bool {
		return history[i].Timestamp.After(history[j].Timestamp)
	})

	if chunk != -1 && chunk < len(history) {
		return history[:chunk], nil
	}
	return history, nil
}
