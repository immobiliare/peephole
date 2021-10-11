package mold

func (db *Mold) Count() (int, error) {
	tx, err := db.Begin(false)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	data, err := tx.GetAll(bucket)
	if err != nil {
		return -1, err
	}

	return len(data), nil
}
