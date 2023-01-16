package database

import (
	"database/sql"
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

func ReadParts() (result []types.Part) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, parts_name, parts_price from parts ")
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {
		var resul types.Part
		err = res.Scan(&resul.Id, &resul.PartsName, &resul.PartsPrice)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}

	return result
}
func NewPart(newPart types.Part) (id int64, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
		return 0, err
	}
	defer db.Close()

	res, err := db.Exec("INSERT INTO `parts` (`parts_name`, `parts_price`) VALUE (?, ?)", newPart.PartsName, newPart.PartsPrice)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()

	return id, err
}
func ChangePart(part types.Part) (err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	_, err = db.Query("UPDATE `parts` SET `parts_name` = ?, `parts_price` = ?  WHERE `id` = ?", part.PartsName, part.PartsPrice, part.Id)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return err
	}

	return err

}
func DeletePart(id int64) (err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("DELETE FROM `parts` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err

}
