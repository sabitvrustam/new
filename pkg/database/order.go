package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sabitvrustam/new/pkg/types"
)

var dbuser string = os.Getenv("bduser")
var dbpass string = os.Getenv("bdpass")
var pass string = fmt.Sprintf("%s:%s@tcp(127.0.0.1)/my_service", dbuser, dbpass)

func ReadOrder(id string) (Order types.Order, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
		return
	}
	defer db.Close()
	var result types.Order
	res, err := db.Query("SELECT o.id, u.f_name, u.l_name, u.m_name, u.n_phone, d.type, d.brand, d.model, d.sn, m.l_name, s.o_status FROM orders AS o "+
		"JOIN users AS u ON o.id_users = u.id "+
		"JOIN device AS d ON o.id_device = d.id "+
		"JOIN masters AS m ON o.id_masters  = m.id "+
		"JOIN status AS s ON o.id_status  = s.id "+
		"WHERE o.id = ?", id)
	if err != nil {
		fmt.Sprintln("не удалось считать данные заказа из базы данных", err)
	}
	for res.Next() {
		err = res.Scan(&result.IdOrder, &result.User.FirstName, &result.User.LastName, &result.User.MidlName, &result.User.Phone, &result.TypeEquipment,
			&result.Brand, &result.Model, &result.Sn, &result.Masters.LastName, &result.Status.StatusOrder)
		if err != nil {
			fmt.Println(err)
		}
	}
	res, err = db.Query("SELECT ow.id_work, w.work_name, w.work_price FROM orders AS o "+
		"JOIN orders_work AS ow ON o.id = ow.id_orders "+
		"JOIN work AS w ON ow.id_work = w.id "+
		"WHERE o.id = ?", id)
	if err != nil {
		fmt.Sprintln("не удалось считать данные работ заказа из базы данных", err)
	}
	for res.Next() {
		var resul types.Work
		err := res.Scan(&resul.Id, &resul.WorkName, &resul.WorkPrice)
		result.Works = append(result.Works, resul)
		if err != nil {
			fmt.Println(err)
		}
	}
	res, err = db.Query("SELECT op.id_parts, p.parts_name , p.parts_price  FROM orders AS o "+
		"JOIN orders_parts AS op ON o.id = op.id_orders "+
		"JOIN parts AS p ON op.id_parts  = p.id "+
		"WHERE o.id = ?", id)
	if err != nil {
		fmt.Sprintln("не удалось считать данные работ заказа из базы данных", err)
	}
	for res.Next() {
		var resul types.Part
		err := res.Scan(&resul.Id, &resul.PartsName, &resul.PartsPrice)
		result.Parts = append(result.Parts, resul)
		if err != nil {
			fmt.Println(err)
		}
	}

	res, err = db.Query("SELECT SUM(parts_price) FROM orders_parts AS op "+
		"JOIN parts AS p ON op.id_parts = p.id "+
		"WHERE id_orders = ?", id)
	if err != nil {
		fmt.Sprintln("не удалось получить сумму запчастей", err)
	}
	for res.Next() {
		err := res.Scan(&result.PriceParts)
		if err != nil {
			fmt.Println("Чтото пошло не так", err)
		}
	}
	res, err = db.Query("SELECT SUM(work_price) FROM orders_work AS ow "+
		"JOIN work AS w ON ow.id_work = w.id "+
		"WHERE id_orders = ?", id)
	if err != nil {
		fmt.Sprintln("не удалось получить сумму запчастей", err)
	}
	for res.Next() {
		err := res.Scan(&result.PriceWork)
		if err != nil {
			fmt.Println("Чтото пошло не так", err)
		}
	}

	return result, err
}
func NewOrder(uw types.Order) (id3 int64, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	res, err := tx.Exec("INSERT INTO `users` (`f_name`, `l_name`, `m_name`, `n_phone`) VALUE (?, ?, ?, ?)", uw.User.FirstName, uw.User.LastName, uw.User.MidlName, uw.User.Phone)
	if err != nil {
		fmt.Println("не удалось записать данные клиента в таблицу", err)
		return 0, err
	}
	id1, err := res.LastInsertId()
	if err != nil {
		fmt.Println("не считать последний добавленный ключ таблицы клиентов", err)
		return 0, err
	}
	fmt.Println(id1)
	res, err = tx.Exec("INSERT INTO `device` (`type`, `brand`, `model`, `sn`) VALUE (?, ?, ?, ?)", uw.TypeEquipment, uw.Brand, uw.Model, uw.Sn)
	if err != nil {
		fmt.Println("не удалось записать данные устроиства в таблицу", err)
		return 0, err
	}
	id2, err := res.LastInsertId()
	if err != nil {
		fmt.Println("не удалось считать последний добавленный ключ таблицы устроиств", err)
		return 0, err
	}
	res, err = tx.Exec("INSERT INTO `orders` (`id_users`, `id_device`, `id_masters`, `id_status`) VALUE (?, ?, ?, ?)", id1, id2, uw.Masters.Id, uw.StatusOrder)
	if err != nil {
		tx.Rollback()
		fmt.Println("не удалось записать ключи в таблицу заказов", err)
		return 0, err
	}
	id3, err = res.LastInsertId()
	err = tx.Commit()
	return id3, err
}
