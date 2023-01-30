package works

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type Work struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewWork(db *sql.DB, log *logrus.Logger) *Work {
	return &Work{
		db:  db,
		log: log}
}

func (d *Work) ReadWoeks() (works []types.Work) {
	sb := sq.Select("id", "work_name", "work_price").From("work")
	res, err := sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err)
	}
	for res.Next() {
		var result types.Work
		err = res.Scan(&result.Id, &result.WorkName, &result.WorkPrice)
		if err != nil {
			d.log.Error(err)
		}
		works = append(works, result)
	}
	return works
}

func (d *Work) WriteWork(newWork types.Work) (id int64, err error) {
	sb := sq.Insert("work").
		Columns("work_name", "work_price").
		Values(newWork.WorkName, newWork.WorkPrice)
	res, err := sb.RunWith(d.db).Exec()
	if err != nil {
		d.log.Error("не удалось записать новую запчасть в базу данных", err)
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		d.log.Error(err)
	}
	return id, err
}

func (d *Work) ChangeWork(work types.Work) (err error) {
	sb := sq.Update("work").
		Set("work_name", work.WorkName).
		Set("work-price", work.WorkPrice).
		Where(sq.Eq{"id": work.Id})
	_, err = sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error("не удалось записать новую запчасть в базу данных", err)
	}
	return err
}

func (d *Work) DeleteWork(id int64) (err error) {
	sb := sq.Delete("work").Where(sq.Eq{"id": id})
	_, err = sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err, "не удалось записать статус ")
	}
	return err

}
