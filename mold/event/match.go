package event

import (
	"regexp"
	"strings"
)

const (
	matchSeparator = "::"
)

func (e *Event) Match(filter string) (match bool) {
	if f := strings.Trim(strings.ToLower(filter), " "); f == "success" || f == "failure" {
		var success bool
		if f == "success" {
			success = true
		}
		match = match || success == e.Success
	}

	r, err := regexp.Compile(filter)
	if err == nil {
		match = match || e.matchReg(r)
	}

	match = match || e.matchStr(filter)
	return
}

func (e *Event) matchGroup() []string {
	return []string{
		e.Jid,
		e.Function,
		e.Minion,
		e.Master,
		strings.Join(e.Args, matchSeparator),
	}
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
	return strings.Contains(strings.Join(e.matchGroup(), matchSeparator), s)
}
