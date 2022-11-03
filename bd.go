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

func dbRead(n int) (result string) {

	db, err := sql.Open("mysql", pass)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("select f_name, l_name, Phone from users where id = ?", n)

	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.FirstName, &user.LastName, &user.Phone)
		if err != nil {
			fmt.Println(err)
		}
		result = fmt.Sprintf("user: %s %s phone nomber %s", user.FirstName, user.LastName, user.Phone)
	}
	return
}

func (m DataRead) dbRead(id string) (User, Equipment) {

	res, err := m.db.Query("select id_users, id_brand from orders where id = ?", id)
	var idOrder Id
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {

		err = res.Scan(&idOrder.IdUser, &idOrder.IdBrands)
		if err != nil {
			fmt.Println(err)
		}

	}
	fmt.Println(idOrder.IdUser, idOrder.IdBrands)

	res, err = m.db.Query("select f_name, l_name, m_name, Phone from users where id = ?", idOrder.IdUser)
	var re User
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {
		err = res.Scan(&re.FirstName, &re.LastName, &re.MidlName, &re.Phone)
		if err != nil {
			fmt.Println(err)
		}
	}

	res, err = m.db.Query("select type, brand, model, SN from brends where id = ?", idOrder.IdBrands)
	var rd Equipment
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {
		err = res.Scan(&rd.TypeEquipment, &rd.Brand, &rd.Model, &rd.Sn)
		if err != nil {
			fmt.Println(err)
		}
	}
	return re, rd

}

func (m DataWrite) dbWrite(uw User, eq Equipment) error {

	tx, err := m.db.Begin()
	if err != nil {
		panic(err)
	}

	result, err := tx.Exec("INSERT INTO `users` (`f_name`, `l_name`, `m_name`, `Phone`) VALUE (?, ?, ?, ?)", uw.FirstName, uw.LastName, uw.MidlName, uw.Phone)
	if err != nil {
		tx.Rollback()
		panic(err)

	}
	id1, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(id1)

	result1, err := tx.Exec("INSERT INTO `brends` (`type`, `brand`, `model`, `SN`) VALUE (?, ?, ?, ?)", eq.TypeEquipment, eq.Brand, eq.Model, eq.Sn)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	id2, err := result1.LastInsertId()
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(id2)

	_, err = tx.Exec("INSERT INTO `orders` (`id_users`, `id_brand`) VALUE (?, ?)", id1, id2)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	err = tx.Commit()
	return err
}
