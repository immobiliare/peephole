package mold

import (
	"sync"

	"github.com/sirupsen/logrus"
	_event "github.com/streambinder/peephole/mold/event"
	"github.com/xujiajun/nutsdb"
)

/*
 * module requirements include:
 * - persistence
 * - auto-expiration
 */

const (
	bucket = "_events"
)

type Mold struct {
	*nutsdb.DB
	config        *Config
	opGetMutex    sync.Mutex
	opGet         chan *_event.Event
	opSelectMutex sync.Mutex
	opSelect      chan []_event.Event
	opCountMutex  sync.Mutex
	opCount       chan int
}

func Init(config *Config) (*Mold, error) {
	opt := nutsdb.DefaultOptions
	opt.Dir = config.Spool
	db, err := nutsdb.Open(opt)
	if err != nil {
		return nil, err
	}

	mold := &Mold{
		db,
		config,
		sync.Mutex{},
		make(chan *_event.Event),
		sync.Mutex{},
		make(chan []_event.Event),
		sync.Mutex{},
		make(chan int),
	}
	go func() {
		logrus.WithField("count", mold.Count()).Println("DB succesfully initialized")
	}()

	return mold, nil
}
