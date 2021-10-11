package event

import (
	"time"
)

type Event struct {
	Master    string
	Minion    string
	Tag       string
	Jid       string
	Raw       string
	Function  string
	Args      []string
	Timestamp time.Time
	Success   bool
}
