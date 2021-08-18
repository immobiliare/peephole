package mold

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	FnHighstate = iota
	FnOrchestrate
	FnSls
	_
	timestampLayout = "2006-01-02T15:04:05.000000"
)

type Event struct {
	Master    string
	Minion    string
	RawData   string
	Function  string
	Timestamp time.Time
	Success   bool
}

func Parse(endpoint, tag, data string) (*Event, error) {
	var (
		jsonData map[string]interface{}
		e        = &Event{Master: endpoint, RawData: data}
	)

	if tag == "salt/auth" {
		return nil, nil
	}

	logrus.Errorln(data)
	if err := json.Unmarshal([]byte(data), &jsonData); err != nil {
		return nil, err
	}

	if _, ok := jsonData["fun"]; !ok {
		return nil, nil
	}
	e.Function = jsonData["fun"].(string)
	if e.Function == "runner.state.orchestrate" {
		return parseOrchestrate(e, jsonData)
	} else {
		return e, nil
	}
}

func parseOrchestrate(e *Event, data map[string]interface{}) (*Event, error) {
	dataFunArgs := data["fun_args"].([]interface{})
	dataPillar := dataFunArgs[0].(map[string]interface{})["pillar"].(map[string]interface{})
	dataEvent := dataPillar["event_data"].(map[string]interface{})

	e.Minion = dataEvent["id"].(string)
	e.Success = dataEvent["retcode"].(float64) == 0

	if t, err := time.Parse(timestampLayout, dataEvent["_stamp"].(string)); err != nil {
		return nil, err
	} else {
		e.Timestamp = t
	}

	return e, nil
}
