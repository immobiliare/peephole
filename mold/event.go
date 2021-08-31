package mold

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

const (
	timestampLayout = "2006-01-02T15:04:05.000000"
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

func Parse(endpoint, tag, data string) (*Event, error) {
	if !gjson.Valid(data) {
		return nil, fmt.Errorf("data structure is not a valid JSON")
	}
	var (
		e = Event{Master: endpoint, Tag: tag, Raw: data}
		j = gjson.Parse(data)
	)

	fun := j.Get("fun").String()
	switch fun {
	case "state.highstate":
		return parseHighstate(&e, &j)
	case "state.sls", "state.apply":
		return parseState(&e, &j)
	case "runner.state.orchestrate":
		return parseOrchestrate(&e, &j)
	}

	return nil, fmt.Errorf("unmanaged type %s", fun)
}

func parseHighstate(e *Event, j *gjson.Result) (*Event, error) {
	e.Function = "highstate"
	e.Args = stringifyResults(j.Get("arg").Array())
	e.Minion = j.Get("id").String()
	e.Success = j.Get("retcode").Int() == 0
	return parseCommon(e, j)
}

func parseState(e *Event, j *gjson.Result) (*Event, error) {
	e.Function = "state"
	e.Args = stringifyResults(j.Get("fun_args").Array())
	e.Minion = j.Get("id").String()
	if j.Get("retcode").Exists() {
		e.Success = j.Get("retcode").Int() == 0
	} else {
		e.Success = j.Get("success").Bool()
	}
	return parseCommon(e, j)
}

func parseOrchestrate(e *Event, j *gjson.Result) (*Event, error) {
	e.Function = "orch"
	e.Args = []string{j.Get("fun_args.0.mods").String()}
	e.Minion = j.Get("fun_args.0.pillar.event_data.id").String()
	e.Success = j.Get("success").Bool()
	return parseCommon(e, j)
}

func parseCommon(e *Event, j *gjson.Result) (*Event, error) {
	e.Jid = j.Get("jid").String()
	if d, err := time.Parse(timestampLayout, j.Get("_stamp").String()); err == nil {
		e.Timestamp = d
	}
	return e, nil
}

func stringifyResults(results []gjson.Result) (arr []string) {
	for _, item := range results {
		arr = append(arr, item.String())
	}
	return
}
