package spy

import (
	"os"

	"github.com/sirupsen/logrus"
	_config "gitlab.rete.farm/dpucci/peephole/config"
	_salt "gitlab.rete.farm/dpucci/peephole/salt-api"
)

type Spy struct {
	endpoints map[string]string
}

func init() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Init(endpoints []*_config.Spy) (*Spy, error) {
	var spy = &Spy{make(map[string]string)}
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
	for k, v := range s.endpoints {
		if err := _salt.Events(k, v); err != nil {
			return err
		}
	}
	return nil
}
