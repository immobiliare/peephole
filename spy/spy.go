package spy

import (
	"os"

	"github.com/sirupsen/logrus"
	_config "gitlab.rete.farm/dpucci/peephole/config"
	_mold "gitlab.rete.farm/dpucci/peephole/mold"
	_salt "gitlab.rete.farm/dpucci/peephole/salt-api"
)

type Spy struct {
	endpoints map[string]string
	EventChan chan *_mold.Event
}

func init() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Init(endpoints []*_config.Spy) (*Spy, error) {
	var spy = &Spy{
		make(map[string]string),
		make(chan *_mold.Event),
	}
	for _, e := range endpoints {
		if r, err := _salt.Login(e.API, e.User, e.Pass, e.Client); err != nil {
			return nil, err
		} else {
			spy.endpoints[e.API] = r.Return[0].Token
			logrus.WithField("token", r.Return[0].Token).Debugln("Token received")
		}
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

		if err := _mold.Persist(o); err != nil {
			logrus.WithError(err).Errorln("Unable to persist event")
			continue
		}

		logrus.WithFields(logrus.Fields{
			"Master":   o.Master,
			"Minion":   o.Minion,
			"Jid":      o.Jid,
			"Function": _mold.FnType[o.Function],
		}).Println("Event persisted, gonna fire it through channel")
		s.EventChan <- o
	}
}
