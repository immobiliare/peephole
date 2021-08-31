package event

import (
	"testing"
)

func TestMatch(t *testing.T) {
	const val = "peephole"
	for k, v := range map[string]Event{
		"jid":      {Jid: val},
		"function": {Function: val},
		"minion":   {Minion: val},
		"master":   {Master: val},
	} {
		if !v.Match(val) {
			t.Errorf("Incorrect unmatch by " + k)
		}
	}
}
