package users

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type Users struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewUser(db *sql.DB, log *logrus.Logger) *Users {
	return &Users{
		db:  db,
		log: log}
}

func (d *Users) ReadUsers(LastName *string, id *int64) (users []types.User, err error) {
	sb := sq.Select("id", "l_name", "f_name", "m_name", "n_phone").
		From("Users")
	if LastName != nil {
		sb = sb.Where(sq.Eq{"l_name": *LastName})
	}
	if id != nil {
		sb = sb.Where(sq.Eq{"id": *id})
	}
	res, err := sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err, "не удалось выполнить запрос селект к таблице мастеров")
	}
	for res.Next() {
		var resul types.User
		err = res.Scan(&resul.Id, &resul.LastName, &resul.FirstName, &resul.MidlName, &resul.Phone)
		if err != nil {
			d.log.Error(err)
		}
		users = append(users, resul)
	}
	return users, err
}

func (d *Users) NewUser1(user types.User) (id int64, err error) {
	sb := sq.Insert("users").
		Columns("l_name", "f_name", "m_name", "n_phone").
		Values(user.FirstName, user.LastName, user.MidlName, user.Phone)
	res, err := sb.RunWith(d.db).Exec()
	if err != nil {
		d.log.Error("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		d.log.Error(err)
	}
	return id, err

}

func (d *Users) ChangUser(user types.User) (err error) {
	sb := sq.Update("users").
		Set("l_name", user.LastName).
		Set("f_name", user.FirstName).
		Set("m_name", user.MidlName).
		Set("n_phone", user.Phone).
		Where(sq.Eq{"id": user.Id})
	_, err = sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error("не удалось записать новую запчасть в базу данных", err)
		return
	}
	return err

}

func (d *Users) DelUser(id int64) (err error) {
	sb := sq.Delete("users").Where(sq.Eq{"id": id})
	_, err = sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err, "не удалось записать статус ")
	}
	return err
}
