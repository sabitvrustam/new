package database

import (
	"database/sql"
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

func ReadOrderPart(id int64) (result []types.OrderParts, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT op.id, op.id_orders, op.id_parts, p.parts_name, p.parts_price from orders_parts AS op "+
		"JOIN parts AS p ON op.id_parts = p.id "+
		"WHERE op.id =?", id)
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице устройств")
	}
	for res.Next() {
		var resul types.OrderParts
		err = res.Scan(&resul.Id, &resul.IdOrder, &resul.IdPart, &resul.PartName, &resul.PartPrice)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func ReadOrderParts(idOrder int64) (result []types.OrderParts, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT op.id, op.id_orders, op.id_parts, p.parts_name, p.parts_price from orders_parts AS op "+
		"JOIN parts AS p ON op.id_parts = p.id "+
		"WHERE op.id_orders =?", idOrder)
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице устройств")
	}
	for res.Next() {
		var resul types.OrderParts
		err = res.Scan(&resul.Id, &resul.IdOrder, &resul.IdPart, &resul.PartName, &resul.PartPrice)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func NewOrderParts(orderParts types.OrderParts) (id int64, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Exec("INSERT INTO `orders_parts` (id_orders, id_parts) VALUE (?, ?)", orderParts.IdOrder,
		orderParts.IdPart)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()

	return id, err

}
func DelOrderParts(id int64) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("DELETE FROM `Orders_parts` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
