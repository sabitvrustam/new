package users

import (
	"database/sql"
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

type Users struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *Users {
	return &Users{db: db}
}

func (d *Users) ReadUsers() (result []types.User, err error) {
	res, err := d.db.Query("SELECT id, l_name, f_name, m_name, n_phone from users ")
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице мастеров")
	}
	for res.Next() {
		var resul types.User
		err = res.Scan(&resul.Id, &resul.LastName, &resul.FirstName, &resul.MidlName, &resul.Phone)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func (d *Users) ReadUsersSearch(LastName string) (result []types.User, err error) {
	res, err := d.db.Query("SELECT id, l_name, f_name, m_name, n_phone from users WHERE f_name = ? ", LastName)
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице мастеров")
	}
	for res.Next() {
		var resul types.User
		err = res.Scan(&resul.Id, &resul.LastName, &resul.FirstName, &resul.MidlName, &resul.Phone)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func (d *Users) ReadUser(id int64) (result []types.User, err error) {
	res, err := d.db.Query("SELECT id, l_name, f_name, m_name, n_phone from users WHERE id = ? ", id)
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице мастеров")
	}
	for res.Next() {
		var resul types.User
		err = res.Scan(&resul.Id, &resul.LastName, &resul.FirstName, &resul.MidlName, &resul.Phone)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func (d *Users) NewUser1(user types.User) (id int64, err error) {
	res, err := d.db.Exec("INSERT INTO `users` (`l_name`, `f_name`, `m_name`, `n_phone` ) VALUE (?, ?, ?, ?)", user.FirstName, user.LastName, user.MidlName, user.Phone)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()
	return id, err

}

func (d *Users) ChangUser(user types.User) (err error) {
	_, err = d.db.Query("UPDATE `users` SET `l_name` = ?, `f_name` = ?, `m_name` = ?, `n_phone` = ? WHERE `id` = ?", user.LastName, user.FirstName, user.MidlName, user.Phone, user.Id)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return err
	}
	return err

}

func (d *Users) DelUser(id int64) (err error) {
	_, err = d.db.Query("DELETE FROM `users` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
