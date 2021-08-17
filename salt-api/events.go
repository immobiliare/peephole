package saltapi

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	_util "gitlab.rete.farm/dpucci/peephole/util"
)

const (
	EventRetry = iota
	EventData
	EventUnknown
)

type EventsResponse struct {
	Type  int
	Tag   string
	Data  string
	Retry int64
}

func Events(endpoint, token string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/events", endpoint), http.NoBody)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Token", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var (
		block   []string
		scanner = bufio.NewScanner(resp.Body)
	)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n ")
		block = append(block, line)
		if line == "" {
			if response, err := unmarshal(block); err == nil {
				// tunnel <- response
				logrus.WithFields(logrus.Fields{
					"type":  response.Type,
					"tag":   response.Tag,
					"data":  response.Data,
					"retry": response.Retry,
				}).Errorln("got a response")
			} else {
				return err
			}
			block = []string{}
		}
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}

func unmarshal(block []string) (*EventsResponse, error) {
	if len(block) == 0 {
		return nil, fmt.Errorf("empty block")
	} else if _util.HasAnyPrefix(block, "retry:") {
		return unmarshalRetry(block)
	} else if _util.HasAnyPrefix(block, "tag:") && _util.HasAnyPrefix(block, "data:") {
		return unmarshalData(block)
	} else {
		return &EventsResponse{Type: EventUnknown}, nil
	}
}

func unmarshalRetry(block []string) (*EventsResponse, error) {
	// not really needed
	r := EventsResponse{Type: EventRetry}
	// for _, msg := range block {
	// 	if strings.HasPrefix("retry:") {
	// 		r.Retry, err = strconv.ParseInt(strings.Split(msg, "retry: ")[1], 10, 64)
	// 	}
	// }
	return &r, nil
}

func unmarshalData(block []string) (*EventsResponse, error) {
	var (
		msgTag  string
		msgData string
	)
	for _, msg := range block {
		if strings.HasPrefix(msg, "tag:") {
			msgTag = strings.Split(msg, "tag: ")[1]
		} else if strings.HasPrefix(msg, "data:") {
			msgData = strings.Split(msg, "data: ")[1]
		}
	}
	return &EventsResponse{Type: EventData, Tag: msgTag, Data: msgData}, nil
}
