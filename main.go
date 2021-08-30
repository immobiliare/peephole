package main

import (
	"flag"
	"os"
	"syscall"

	"github.com/sirupsen/logrus"
	_config "github.com/streambinder/peephole/config"
	_kiosk "github.com/streambinder/peephole/kiosk"
	_mold "github.com/streambinder/peephole/mold"
	_spy "github.com/streambinder/peephole/spy"
	_util "github.com/streambinder/peephole/util"
)

var (
	argConfig string

	config *_config.Wrapper
	mold   *_mold.Mold
	spy    *_spy.Spy
	kiosk  *_kiosk.Kiosk
)

func init() {
	_util.TrapSignal(os.Interrupt, func() { exit() })
	flag.StringVar(&argConfig, "c", "/etc/peephole", "Configuration file path")
	flag.Parse()
}

func main() {
	var err error
	logrus.WithField("path", argConfig).Infoln("Reading configuration file")
	if config, err = _config.Parse(argConfig); err != nil {
		logrus.WithError(err).Fatalln("Unable to get config")
	}

	if config.Debug && !_util.Debugging() {
		if err := syscall.Exec(os.Args[0], os.Args,
			append(os.Environ(), []string{_util.DebugKey}...)); err != nil {
			logrus.WithError(err).Fatalln("Unable to exec proc in debug mode")
		}
		exit()
	}

	logrus.Infoln("Initializing mold")
	if mold, err = _mold.Init(config.Mold); err != nil {
		logrus.WithError(err).Fatalln("Unable to get config")
	}

	logrus.Infoln("Initializing spy")
	if spy, err = _spy.Init(mold, config.Spy); err != nil {
		logrus.WithError(err).Errorln("Unable to get config")
		exit()
	}

	logrus.WithField("endpoints", len(spy.Endpoints())).Infoln("Starting spying")
	go func() {
		if err := spy.Watch(); err != nil {
			logrus.WithError(err).Errorln("Unable to spy")
			exit()
		}
	}()

	logrus.Infoln("Starting kiosk server")
	kiosk = _kiosk.Init(mold, spy.EventChan, config.Kiosk)
	if err := kiosk.Serve(); err != nil {
		logrus.WithError(err).Errorln("Unable to serve connections")
		exit()
	}
}

func exit() {
	if mold == nil {
		return
	}

	logrus.Println("Closing DB")
	if err := mold.Close(); err != nil {
		logrus.WithError(err).Println("Unable to close db")
	}

	os.Exit(0)
}
