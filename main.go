package main

import (
	"flag"

	"github.com/sirupsen/logrus"
	_config "gitlab.rete.farm/dpucci/peephole/config"
	_kiosk "gitlab.rete.farm/dpucci/peephole/kiosk"
	_spy "gitlab.rete.farm/dpucci/peephole/spy"
)

var (
	argConfig string

	config *_config.Wrapper
	spy    *_spy.Spy
	kiosk  *_kiosk.Kiosk
)

func init() {
	flag.StringVar(&argConfig, "c", "/etc/peephole", "Configuration file path")
	flag.Parse()
}

func main() {
	var err error
	logrus.WithField("path", argConfig).Infoln("Reading configuration file")
	if config, err = _config.Parse(argConfig); err != nil {
		logrus.WithError(err).Fatalln("Unable to get config")
	}

	logrus.Infoln("Initializing spy")
	if spy, err = _spy.Init(config.Spy); err != nil {
		logrus.WithError(err).Fatalln("Unable to get config")
	}

	logrus.WithField("endpoints", len(spy.Endpoints())).Infoln("Starting spying")
	go func() {
		if err := spy.Watch(); err != nil {
			logrus.WithError(err).Fatalln("Unable to spy")
		}
	}()

	logrus.Infoln("Starting kiosk server")
	kiosk = _kiosk.Init(spy.EventChan, config.Kiosk)
	if err := kiosk.Serve(); err != nil {
		logrus.WithError(err).Errorln("Unable to serve connections")
	}
}
