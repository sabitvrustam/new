package database

import (
	"database/sql"
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

func ReadMasters() (result []types.Master, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, l_name, f_name, m_name, n_phone from masters ")
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице мастеров")
	}
	for res.Next() {
		var resul types.Master
		err = res.Scan(&resul.Id, &resul.LastName, &resul.FirstName, &resul.MidlName, &resul.Phone)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func NewMaster(master types.Master) (id int64, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Exec("INSERT INTO `masters` (`l_name`, `f_name`, `m_name`, `n_phone` ) VALUE (?, ?, ?, ?)", master.FirstName, master.LastName, master.MidlName, master.Phone)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()

	return id, err

}
func ChangMaster(master types.Master) (err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	_, err = db.Query("UPDATE `masters` SET `l_name` = ?, `f_name` = ?, `m_name` = ?, `n_phone` = ? WHERE `id` = ?", master.LastName, master.FirstName, master.MidlName, master.Phone, master.Id)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return err
	}

	return err

}
func DeleteMaster(id int64) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("DELETE FROM `masters` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
