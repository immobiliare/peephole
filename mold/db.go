package mold

var (
	history []*Event
)

func Persist(e *Event) error {
	history = append(history, e)
	return nil
}

func Select() ([]*Event, error) {
	return history, nil
}
