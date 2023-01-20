package database

import (
	"database/sql"
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

func ReadUsers() (result []types.User, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, l_name, f_name, m_name, n_phone from users ")
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

func ReadUsersSearch(LastName string) (result []types.User, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, l_name, f_name, m_name, n_phone from users WHERE f_name = ? ", LastName)
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
func ReadUser(id int64) (result []types.User, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, l_name, f_name, m_name, n_phone from users WHERE id = ? ", id)
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

func NewUser(user types.User) (id int64, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Exec("INSERT INTO `users` (`l_name`, `f_name`, `m_name`, `n_phone` ) VALUE (?, ?, ?, ?)", user.FirstName, user.LastName, user.MidlName, user.Phone)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()

	return id, err

}
func ChangUser(user types.User) (err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы клиентов", err)
	}
	defer db.Close()
	_, err = db.Query("UPDATE `users` SET `l_name` = ?, `f_name` = ?, `m_name` = ?, `n_phone` = ? WHERE `id` = ?", user.LastName, user.FirstName, user.MidlName, user.Phone, user.Id)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return err
	}

	return err

}
func DeleteUser(id int64) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("DELETE FROM `users` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
