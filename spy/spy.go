package spy

import (
	"strings"
	"time"

	_mold "github.com/immobiliare/peephole/mold"
	_event "github.com/immobiliare/peephole/mold/event"
	_salt "github.com/immobiliare/peephole/salt"
	_util "github.com/immobiliare/peephole/util"
	"github.com/sirupsen/logrus"
)

type Spy struct {
	endpoints map[string]*endpoint
	EventChan chan *_event.Event
	db        *_mold.Mold
}

type endpoint struct {
	user     string
	password string
	client   string
	token    string
}

func init() {
	if _util.Debugging() {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Init(db *_mold.Mold, endpoints []*Config) (*Spy, error) {
	var spy = &Spy{
		make(map[string]*endpoint),
		make(chan *_event.Event),
		db,
	}
	for _, e := range endpoints {
		endpoint := &endpoint{e.User, e.Pass, e.Client, ""}
		token, err := login(e.API, endpoint)
		if err != nil {
			return nil, err
		}

		endpoint.token = token
		spy.endpoints[e.API] = endpoint
		logrus.WithField("token", token).Debugln("Token received")
	}
	return spy, nil
}

func (s *Spy) Endpoints() (e []string) {
	for k := range s.endpoints {
		e = append(e, k)
	}
	return
}

func (s *Spy) Watch() error {
	peephole := make(chan *_salt.EventsResponse, len(s.endpoints))
	for k, v := range s.endpoints {
		go func(api string, endpoint *endpoint) {
			if err := s.spy(api, endpoint, peephole); err != nil {
				logrus.WithError(err).WithField("endpoint", endpoint).Fatalln("Unable to spy")
			}
		}(k, v)
	}

	for {
		e := <-peephole
		logrus.WithFields(logrus.Fields{
			"endpoint": e.Endpoint,
			"tag":      e.Tag,
		}).Debugln("Event received")

		o, err := _event.Parse(e.Endpoint, e.Tag, e.Data)
		if err != nil {
			logrus.WithError(err).Warnln("Unable to parse event")
			continue
		}

		if err := s.db.Write(o); err != nil {
			logrus.WithError(err).Errorln("Unable to persist event")
			continue
		}

		logrus.WithFields(logrus.Fields{
			"master":   o.Master,
			"minion":   o.Minion,
			"jid":      o.Jid,
			"function": o.Function,
			"args":     strings.Join(o.Args, ","),
		}).Println("Event persisted")
		s.EventChan <- o
	}
}

func (s *Spy) spy(api string, endpoint *endpoint, peephole chan *_salt.EventsResponse) error {
	for {
		if s.endpoints[api].token == "EXPIRED" {
			logrus.WithField("endpoint", api).Println("Reauthenticating")
			tok, err := login(api, endpoint)
			if err != nil {
				logrus.WithError(err).WithField("endpoint", api).Errorln("Unable to authenticate")
				continue
			}
			s.endpoints[api].token = tok
		}

		err := _salt.Events(api, endpoint.token, peephole)
		if err != nil {
			logrus.WithError(err).WithField("endpoint", api).Warnln("Spying interrupted, retrying")
			s.endpoints[api].token = "EXPIRED"
		}

		time.Sleep(5 * time.Second)
	}
}

func login(api string, endpoint *endpoint) (string, error) {
	r, err := _salt.Login(api, endpoint.user, endpoint.password, endpoint.client)
	if err != nil {
		return "", err
	}
	return r.Return[0].Token, nil
}
