package status

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type Status struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewStatus(db *sql.DB, log *logrus.Logger) *Status {
	return &Status{
		db:  db,
		log: log}

}

func (d *Status) ReadStatus() (status []types.Status) {

	sb := sq.Select("id", "o_status").From("status")
	res, err := sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err)
	}
	for res.Next() {
		var result types.Status
		err = res.Scan(&result.Id, &result.StatusOrder)
		if err != nil {
			d.log.Error(err)
		}
		status = append(status, result)
	}
	return status
}
