package salt

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	_util "github.com/immobiliare/peephole/util"
	"github.com/sirupsen/logrus"
)

const (
	eventRetry = iota
	eventData
	eventUnknown
)

var (
	eventTypes = map[int]string{
		eventRetry:   "Retry",
		eventData:    "Data",
		eventUnknown: "Unknown",
	}
	eventTagRegex = regexp.MustCompile(`(?m)(salt/run/[0-9]{20}/ret|salt/job/[0-9]{20}/ret/.*)`)
)

type EventsResponse struct {
	Endpoint string
	Type     int
	Tag      string
	Data     string
	Retry    int64
}

func init() {
	if _util.Debugging() {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Events(endpoint, token string, peephole chan *EventsResponse) error {
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
	// max batch size: 32, max message size: 4MB
	scanner.Buffer(make([]byte, 1024*1024), 128*1024*1024)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n ")
		if strings.Contains(line, "401 Unauthorized") {
			return fmt.Errorf("unauthorized: invalid token")
		}

		block = append(block, line)
		if line == "" {
			if response, err := unmarshal(block); err == nil {
				response.Endpoint = endpoint
				if eventTagRegex.Match([]byte(response.Tag)) {
					logrus.WithFields(logrus.Fields{
						"type":     eventTypes[response.Type],
						"endpoint": response.Endpoint,
						"tag":      response.Tag,
						"data":     response.Data,
						// "retry": response.Retry,
					}).Debugln("Event received")
					peephole <- response
				}
			} else {
				return err
			}
			block = []string{}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func unmarshal(block []string) (*EventsResponse, error) {
	if _util.HasAnyPrefix(block, "retry:") {
		return unmarshalRetry(block)
	} else if _util.HasAnyPrefix(block, "tag:") && _util.HasAnyPrefix(block, "data:") {
		return unmarshalData(block)
	}
	return &EventsResponse{Type: eventUnknown}, nil
}

func unmarshalRetry(block []string) (*EventsResponse, error) {
	// not really needed
	// let's just mark the message per type
	r := EventsResponse{Type: eventRetry}
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
			// tag prefix drop
			msgTag = strings.Split(msg, "tag: ")[1]
		} else if strings.HasPrefix(msg, "data:") {
			// tag/data prefixes drop
			// sample line expected:
			// `data: {"tag": "whatevert-tag", "data": { ... }}``
			tagField := fmt.Sprintf(`"tag": "%s"`, msgTag)
			msgData = msg[6:]                                                        // data: prefix
			msgData = strings.Replace(msgData, fmt.Sprintf("%s, ", tagField), "", 1) // tag as a heading field
			msgData = strings.Replace(msgData, fmt.Sprintf(", %s", tagField), "", 1) // tag as a tailing field
			msgData = strings.Replace(msgData, `{"data": `, "", 1)                   // data header
			// trailing curly bracket
			msgData = msgData[:len(msgData)-1]
		}
	}
	return &EventsResponse{Type: eventData, Tag: msgTag, Data: msgData}, nil
}
