package mold

import (
	"github.com/sirupsen/logrus"
)

func (db *Mold) Count(filter string) (int, error) {
	var (
		data [][]byte
		err  error
	)

	tx, err := db.Begin(false)
	if err != nil {
		return -1, err
	}
	defer func() {
		err := tx.Rollback()
		if err != nil {
			logrus.Debugln(err.Error())
		}
	}()

	if filter != "" {
		data, err = tx.PrefixSearchScan(bucket, []byte{}, filter, 0, -1)
	} else {
		data, err = tx.PrefixScan(bucket, []byte{}, 0, -1)
	}

	if err != nil &&
		(err.Error() == "prefix scans not found" || err.Error() == "prefix and search scans not found") {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return len(data), nil
}
