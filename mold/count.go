package mold

import (
	"github.com/sirupsen/logrus"
	"github.com/xujiajun/nutsdb"
)

func (db *Mold) Count() int {
	db.opCountMutex.Lock()
	defer db.opCountMutex.Unlock()

	go func() {
		if err := db.View(
			func(tx *nutsdb.Tx) error {
				data, err := tx.GetAll(bucket)
				if err != nil {
					db.opCount <- 0
					return err
				}

				db.opCount <- len(data)
				return nil
			}); err != nil {
			logrus.WithError(err).Errorln("Unable to count keys")
		}
	}()

	return <-db.opCount
}
