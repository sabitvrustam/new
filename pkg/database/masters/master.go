package masters

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type Master struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewMaster(db *sql.DB, log *logrus.Logger) *Master {
	return &Master{
		db:  db,
		log: log}
}

func (d *Master) MastersRead() (result []types.Master, err error) {
	masters := sq.Select("id", "l_name", "f_name", "m_name").
		From("masters")
	res, err := masters.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err, "не удалось выполнить запрос селект к таблице мастеров")
	}
	for res.Next() {
		var resul types.Master
		err = res.Scan(&resul.Id, &resul.LastName, &resul.FirstName, &resul.MidlName, &resul.Phone)
		if err != nil {
			d.log.Error(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func (d *Master) MastersWrite(result types.Master) (id int64, err error) {
	masters := sq.Insert("masters").
		Columns("l_name", "f_name", "m_name", "n_phone").
		Values(result.FirstName, result.LastName, result.MidlName, result.Phone)
	res, err := masters.RunWith(d.db).Exec()
	if err != nil {
		d.log.Error("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		d.log.Error("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	return id, err
}

func (d *Master) MastersChange(result types.Master) (err error) {
	masters := sq.Update("masters").
		Set("l_name", result.LastName).
		Set("f_name", result.LastName).
		Set("m_name", result.MidlName).
		Set("n_phone", result.Phone)
	_, err = masters.RunWith(d.db).Exec()
	if err != nil {
		d.log.Error("не удалось записать новую запчасть в базу данных", err)
	}
	return err
}

func (d *Master) MastersDelete(id int64) (err error) {
	masters := sq.Delete("masters").Where(sq.Eq{"id": id})
	_, err = masters.RunWith(d.db).Exec()
	if err != nil {
		d.log.Error(err, "не удалось записать статус ")
	}
	return err
}
