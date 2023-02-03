package count

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Count struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewCount(db *sql.DB, log *logrus.Logger) *Count {
	return &Count{
		db:  db,
		log: log,
	}
}

func (d *Count) CountRead() (count int64, err error) {
	res, err := d.db.Query("select count(*) from `orders`")
	for res.Next() {
		err = res.Scan(&count)
		if err != nil {
			d.log.Error(err)
		}
	}
	return count, err
}
