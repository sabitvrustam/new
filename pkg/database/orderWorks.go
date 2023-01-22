package database

import (
	"database/sql"
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

func ReadOrderWorks(id int64) (result []types.OrderWorks, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT ow.id, ow.id_orders, ow.id_work, w.work_name, w.work_price from orders_work AS ow "+
		"JOIN work AS w ON ow.id_work = w.id "+
		"WHERE ow.id_orders =?", id)
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице устройств")
	}
	for res.Next() {
		var resul types.OrderWorks
		err = res.Scan(&resul.Id, &resul.IdOrder, &resul.IdWork, &resul.WorkName, &resul.WorkPrice)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}
func ReadOrderWork(id int64) (result []types.OrderWorks, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT ow.id, ow.id_orders, ow.id_work, w.work_name, w.work_price from orders_work AS ow "+
		"JOIN work AS w ON ow.id_work = w.id "+
		"WHERE ow.id =?", id)
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице устройств")
	}
	for res.Next() {
		var resul types.OrderWorks
		err = res.Scan(&resul.Id, &resul.IdOrder, &resul.IdWork, &resul.WorkName, &resul.WorkPrice)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func NewOrderWorks(orderWorks types.OrderWorks) (id int64, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Exec("INSERT INTO `orders_work` (id_orders, id_work) VALUE (?, ?)", orderWorks.IdOrder,
		orderWorks.IdWork)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()

	return id, err

}
func DelOrderWorks(id int64) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("DELETE FROM `orders_work` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
