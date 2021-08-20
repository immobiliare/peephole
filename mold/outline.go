package mold

func (e *Event) Outline() Event {
	return Event{
		Master:    e.Master,
		Minion:    e.Minion,
		Tag:       e.Tag,
		Jid:       e.Jid,
		Function:  e.Function,
		Timestamp: e.Timestamp,
		Success:   e.Success,
	}
}
