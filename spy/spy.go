package spy

import (
	"os"

	"github.com/sirupsen/logrus"
	_mold "github.com/streambinder/peephole/mold"
	_salt "github.com/streambinder/peephole/salt"
)

type Spy struct {
	endpoints map[string]string
	EventChan chan *_mold.Event
	db        *_mold.Mold
}

func init() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Init(db *_mold.Mold, endpoints []*Config) (*Spy, error) {
	var spy = &Spy{
		make(map[string]string),
		make(chan *_mold.Event),
		db,
	}
	for _, e := range endpoints {
		r, err := _salt.Login(e.API, e.User, e.Pass, e.Client)
		if err != nil {
			return nil, err
		}

		spy.endpoints[e.API] = r.Return[0].Token
		logrus.WithField("token", r.Return[0].Token).Debugln("Token received")
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
		go func(endpoint, token string, peephole chan *_salt.EventsResponse) {
			if err := _salt.Events(endpoint, token, peephole); err != nil {
				logrus.WithError(err).Fatalln("Unable to watch for events")
			}
		}(k, v, peephole)
	}

	for {
		e := <-peephole
		logrus.WithFields(logrus.Fields{
			"Endpoint": e.Endpoint,
			"Tag":      e.Tag,
		}).Debugln("Event received")

		o, err := _mold.Parse(e.Endpoint, e.Tag, e.Data)
		if err != nil {
			logrus.WithError(err).Warnln("Unable to parse event")
			continue
		}

		if err := s.db.Write(o); err != nil {
			logrus.WithError(err).Errorln("Unable to persist event")
			continue
		}

		logrus.WithFields(logrus.Fields{
			"Master":   o.Master,
			"Minion":   o.Minion,
			"Jid":      o.Jid,
			"Function": o.Function,
		}).Println("Event persisted")
		s.EventChan <- o
	}
}
