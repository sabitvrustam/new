package database

import (
	"database/sql"
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

func ReadWoeks() (result []types.Work) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с списка работ", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, work_name, work_price from work ")
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {
		var resul types.Work
		err = res.Scan(&resul.Id, &resul.WorkName, &resul.WorkPrice)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}

	return result
}

func WriteWork(newWork types.Work) (id int64, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
		return 0, err
	}
	defer db.Close()

	res, err := db.Exec("INSERT INTO `work` (`work_name`, `work_price`) VALUE (?, ?)", newWork.WorkName, newWork.WorkPrice)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()
	return id, err
}

func ChangeWork(work types.Work) (err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	_, err = db.Query("UPDATE `work` SET `work_name` = ?, `work_price` = ?  WHERE `id` = ?", work.WorkName, work.WorkPrice, work.Id)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return err
	}

	return err

}
func DeleteWork(id int64) (err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("DELETE FROM `work` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err

}
