package mold

import (
	"fmt"
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
	RawData   string
	Function  string
	Timestamp time.Time
	Success   bool
}

func Parse(endpoint, tag, data string) (*Event, error) {
	if !gjson.Valid(data) {
		return nil, fmt.Errorf("data structure is not a valid JSON")
	}
	var (
		e = Event{Master: endpoint, Tag: tag, RawData: data}
		j = gjson.Parse(data)
	)

	fun := j.Get("fun").String()
	if fun == "state.highstate" {
		return parseHighstate(&e, &j)
	} else if fun == "state.sls" || fun == "state.apply" {
		return parseState(&e, &j)
	} else if fun == "runner.state.orchestrate" {
		return parseOrchestrate(&e, &j)
	}

	return nil, fmt.Errorf("unmanaged type %s", fun)
}

func parseHighstate(e *Event, j *gjson.Result) (*Event, error) {
	e.Function = "highstate"
	e.Minion = j.Get("id").String()
	e.Jid = j.Get("jid").String()
	e.Success = j.Get("retcode").Int() == 0
	if d, err := time.Parse(timestampLayout, j.Get("_stamp").String()); err == nil {
		e.Timestamp = d
	}
	return e, nil
}

func parseState(e *Event, j *gjson.Result) (*Event, error) {
	e.Function = "state apply"
	e.Minion = j.Get("id").String()
	e.Jid = j.Get("jid").String()
	e.Success = j.Get("success").Bool()
	if d, err := time.Parse(timestampLayout, j.Get("_stamp").String()); err == nil {
		e.Timestamp = d
	}
	return e, nil
}

func parseOrchestrate(e *Event, j *gjson.Result) (*Event, error) {
	e.Function = "orchestrate"
	e.Minion = j.Get("fun_args.0.pillar.event_data.id").String()
	e.Jid = j.Get("jid").String()
	e.Success = j.Get("success").Bool()
	if d, err := time.Parse(timestampLayout, j.Get("_stamp").String()); err == nil {
		e.Timestamp = d
	}
	return e, nil
}
