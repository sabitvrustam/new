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
		var user UserRead
		err = res.Scan(&user.FirstName, &user.LastName, &user.Phone)
		if err != nil {
			fmt.Println(err)
		}
		result = fmt.Sprintf("user: %s %s phone nomber %s", user.FirstName, user.LastName, user.Phone)
	}
	return
}

func (m DataRead) dbRead(id string) UserRead {

	res, err := m.db.Query("select f_name, l_name, Phone from users where id = ?", id)
	var re UserRead
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {

		err = res.Scan(&re.FirstName, &re.LastName, &re.Phone)
		if err != nil {
			fmt.Println(err)
		}

	}
	return re
}

func (m DataWrite) dbWrite(uw UserWrite, eq Equipment) error {

	tx, err := m.db.Begin()
	if err != nil {
		panic(err)
	}

	result, err := tx.Exec("INSERT INTO `users` (`f_name`, `l_name`, `Phone`) VALUE (?, ?, ?)", uw.firstName, uw.lastName, uw.phone)
	if err != nil {
		tx.Rollback()
		panic(err)

	}
	id1, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(id1)

	result1, err := tx.Exec("INSERT INTO `brends` (`type`, `brand`, `model`, `SN`) VALUE (?, ?, ?, ?)", eq.typeEquipment, eq.brand, eq.model, eq.sn)
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
