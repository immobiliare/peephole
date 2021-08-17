package main

import (
	"flag"

	"github.com/sirupsen/logrus"
	_config "gitlab.rete.farm/dpucci/peephole/config"
	_spy "gitlab.rete.farm/dpucci/peephole/spy"
)

var (
	argConfig string

	config *_config.Wrapper
	spy    *_spy.Spy
)

func init() {
	flag.StringVar(&argConfig, "c", "/etc/peephole", "Configuration file path")
	flag.Parse()
}

func main() {
	var err error
	logrus.WithField("path", argConfig).Infoln("Reading configuration file")
	if config, err = _config.Parse(argConfig); err != nil {
		logrus.WithError(err).Fatalln("unable to get config")
	}

	logrus.Infoln("Initializing spy")
	if spy, err = _spy.Init(config.Spy); err != nil {
		logrus.WithError(err).Fatalln("unable to get config")
	}

	logrus.WithField("endpoints", len(spy.Endpoints())).Infoln("Starting spying")
	if err := spy.Watch(); err != nil {
		logrus.WithError(err).Fatalln("unable to get config")
	}
}
