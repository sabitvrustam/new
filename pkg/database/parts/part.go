package parts

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type Part struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewPart(db *sql.DB, log *logrus.Logger) *Part {
	return &Part{
		db:  db,
		log: log}
}

func (d *Part) ReadParts() (parts []types.Part) {
	sb := sq.Select("id", "parts_name", "parts_price").From("parts")
	res, err := sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err)
	}
	for res.Next() {
		var result types.Part
		err = res.Scan(&result.Id, &result.PartsName, &result.PartsPrice)
		if err != nil {
			d.log.Error(err)
		}
		parts = append(parts, result)
	}
	return parts
}

func (d *Part) NewPart1(newPart types.Part) (id int64, err error) {
	sb := sq.Insert("parts").
		Columns("parts_name", "parts_price").
		Values(newPart.PartsName, newPart.PartsPrice)
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

func (d *Part) ChangePart(part types.Part) (err error) {
	sb := sq.Update("parts").
		Set("parts_name", part.PartsName).
		Set("parts_price", part.PartsPrice).
		Where(sq.Eq{"id": part.Id})
	_, err = sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error("не удалось записать новую запчасть в базу данных", err)
		return
	}
	return err
}

func (d *Part) DeletePart(id int64) (err error) {
	sb := sq.Delete("parts").Where(sq.Eq{"id": id})
	_, err = sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err, "не удалось записать статус ")
	}
	return err
}
