package event

import (
	"fmt"
	"testing"
	"time"
)

var (
	master    = "master"
	tag       = "tag"
	minion    = "minion"
	success   = true
	jid       = "0"
	timestamp = "2000-01-01T10:00:00.000000"
	retvals   = map[bool]int{
		true:  0,
		false: 1,
	}
)

func TestParseHighstate(t *testing.T) {
	var (
		fun  = "state.highstate"
		json = fmt.Sprintf(`{
		"fun": "%s",
		"id": "%s",
		"retcode": %d,
		"jid": "%s",
		"_stamp": "%s"
	}`, fun, minion, retvals[success], jid, timestamp)
	)

	e, err := Parse(master, tag, json)
	if err != nil {
		t.Errorf("Unexpected error while parsing json data")
	}

	if !(e.Function[:9] == "highstate" && condCommon(e)) {
		t.Errorf("Parsed event mismatches with input data")
	}
}

func TestParseStateSLS(t *testing.T) {
	for _, fun := range []string{"state.sls", "state.apply"} {
		var json = fmt.Sprintf(`{
			"fun": "%s",
			"id": "%s",
			"success": %t,
			"jid": "%s",
			"_stamp": "%s"
		}`, fun, minion, success, jid, timestamp)

		e, err := Parse(master, tag, json)
		if err != nil {
			t.Errorf("Unexpected error while parsing json data")
		}

		if !(e.Function[:5] == "state" && condCommon(e)) {
			t.Errorf("Parsed event mismatches with input data")
		}
	}
}

func TestParseOrchestrate(t *testing.T) {
	var (
		fun  = "runner.state.orchestrate"
		json = fmt.Sprintf(`{
		"fun": "%s",
		"fun_args": [{"pillar": {"event_data": {"id": "%s"}}}],
		"success": %t,
		"jid": "%s",
		"_stamp": "%s"
	}`, fun, minion, success, jid, timestamp)
	)

	e, err := Parse(master, tag, json)
	if err != nil {
		t.Errorf("Unexpected error while parsing json data")
	}

	if !(e.Function[:4] == "orch" && condCommon(e)) {
		t.Errorf("Parsed event mismatches with input data")
	}
}

func condCommon(e *Event) (cond bool) {
	cond = e.Master == master
	cond = cond && e.Tag == tag
	cond = cond && e.Minion == minion
	cond = cond && e.Success == success
	cond = cond && e.Jid == jid
	ts, err := time.Parse(timestampLayout, timestamp)
	if err != nil {
		return false
	}
	cond = cond && e.Timestamp == ts
	return
}
