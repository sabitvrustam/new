package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbuser string = os.Getenv("bduser")
var dbpass string = os.Getenv("bdpass")
var pass string = fmt.Sprintf("%s:%s@tcp(127.0.0.1)/my_service", dbuser, dbpass)

func dbRead(id string) Order {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	defer db.Close()
	var result Order
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
		err = res.Scan(&result.IdOrder, &result.FirstName, &result.LastName, &result.MidlName, &result.Phone, &result.TypeEquipment,
			&result.Brand, &result.Model, &result.Sn, &result.Masters.L_name, &result.Status.StatusOrder)
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
		var resul Work
		err := res.Scan(&resul.IdWork, &resul.WorkName, &resul.WorkPrice)
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
		var resul Part
		err := res.Scan(&resul.IdPart, &resul.PartsName, &resul.PartsPrice)
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

	return result
}
func dbWritePartsOrder(in Order) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("INSERT INTO `orders_parts` (`id_orders`, `id_parts`) VALUE (?,?)", in.IdOrder, in.Part.Id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
func dbDeletePartsOrder(idPart Order) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("DELETE FROM `orders_parts` WHERE `id_orders`=? AND `id_parts`=?", idPart.IdOrder, idPart.Part.Id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
func dbWriteWorksOrder(in Order) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("INSERT INTO `orders_work` (`id_orders`, `id_work`) VALUE (?,?)", in.IdOrder, in.Work.Id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
func dbDeleteWorksOrder(idwork Order) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("DELETE FROM `orders_work` WHERE `id_orders`=? AND `id_work`=?", idwork.IdOrder, idwork.Work.Id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
func dbWrite(uw Order) (err error, id3 int64) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	res, err := tx.Exec("INSERT INTO `users` (`f_name`, `l_name`, `m_name`, `n_phone`) VALUE (?, ?, ?, ?)", uw.FirstName, uw.LastName, uw.MidlName, uw.Phone)
	if err != nil {
		fmt.Println("не удалось записать данные клиента в таблицу", err)
		return err, 0
	}
	id1, err := res.LastInsertId()
	if err != nil {
		fmt.Println("не считать последний добавленный ключ таблицы клиентов", err)
		return err, 0
	}
	fmt.Println(id1)
	res, err = tx.Exec("INSERT INTO `device` (`type`, `brand`, `model`, `sn`) VALUE (?, ?, ?, ?)", uw.TypeEquipment, uw.Brand, uw.Model, uw.Sn)
	if err != nil {
		fmt.Println("не удалось записать данные устроиства в таблицу", err)
		return err, 0
	}
	id2, err := res.LastInsertId()
	if err != nil {
		fmt.Println("не удалось считать последний добавленный ключ таблицы устроиств", err)
		return err, 0
	}
	fmt.Println(id2, uw.Masters.Id)
	res, err = tx.Exec("INSERT INTO `orders` (`id_users`, `id_device`, `id_masters`, `id_status`) VALUE (?, ?, ?, ?)", id1, id2, uw.Masters.Id, uw.StatusOrder)
	if err != nil {
		tx.Rollback()
		fmt.Println("не удалось записать ключи в таблицу заказов", err)
		return err, 0
	}
	id3, err = res.LastInsertId()
	err = tx.Commit()
	return err, id3
}
func dbreadMasters() []Masters {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, l_name from masters ")
	var result []Masters
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {
		var resul Masters
		err = res.Scan(&resul.Id, &resul.L_name)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result
}

func dbreadParts() Order {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, parts_name, parts_price from parts ")
	var result Order
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {
		var resul Part
		err = res.Scan(&resul.Id, &resul.PartsName, &resul.PartsPrice)
		if err != nil {
			fmt.Println(err)
		}
		result.OllParts = append(result.OllParts, resul)
	}

	return result
}
func dbreadWorks() Order {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с списка работ", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, work_name, work_price from work ")
	var result Order
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {
		var resul Work
		err = res.Scan(&resul.Id, &resul.WorkName, &resul.WorkPrice)
		if err != nil {
			fmt.Println(err)
		}
		result.OllWorks = append(result.OllWorks, resul)
	}

	return result
}

func dbWriteParts(newPart Part) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
		return err
	}
	defer db.Close()

	_, err = db.Query("INSERT INTO `parts` (`parts_name`, `parts_price`) VALUE (?, ?)", newPart.PartsName, newPart.PartsPrice)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return err
	}

	return err
}
func dbWriteWork(newWork Work) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
		return err
	}
	defer db.Close()

	_, err = db.Query("INSERT INTO `work` (`work_name`, `work_price`) VALUE (?, ?)", newWork.WorkName, newWork.WorkPrice)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return err
	}

	return err
}

// func dbreadStatus() []Status {
// 	db, err := sql.Open("mysql", pass)
// 	if err != nil {
// 		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
// 	}
// 	defer db.Close()
// 	res, err := db.Query("SELECT id, status from status ")
// 	var result []Status
// 	if err != nil {
// 		fmt.Sprintln(err)
// 	}
// 	for res.Next() {
// 		var resul Masters
// 		err = res.Scan(&resul.Id, &resul.L_name)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		result = append(result, resul)
// 	}
// 	fmt.Println(result)
// 	return result
// }
