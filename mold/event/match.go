package event

import (
	"regexp"
	"strings"
)

func (e *Event) Match(filter string) bool {
	r, err := regexp.Compile(filter)
	if err == nil {
		return e.matchReg(r)
	}
	return e.matchStr(filter)
}

func (e *Event) matchGroup() []string {
	return []string{e.Jid, e.Function, e.Minion, e.Master}
}

func (e *Event) matchReg(r *regexp.Regexp) bool {
	for _, k := range e.matchGroup() {
		if r.Match([]byte(k)) {
			return true
		}
	}
	return false
}

func (e *Event) matchStr(s string) bool {
	return strings.Contains(strings.Join(e.matchGroup(), "::"), s)
}
