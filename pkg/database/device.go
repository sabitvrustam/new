package database

import (
	"database/sql"
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

func ReadDevices() (result []types.Device, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, type, brand, model, sn from device ")
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице устройств")
	}
	for res.Next() {
		var resul types.Device
		err = res.Scan(&resul.Id, &resul.TypeEquipment, &resul.Brand, &resul.Model, &resul.Sn)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func ReadDevicesSearch(sn string) (result []types.Device, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, type, brand, model, sn from device WHERE sn = ? ", sn)
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице мастеров")
	}
	for res.Next() {
		var resul types.Device
		err = res.Scan(&resul.Id, &resul.TypeEquipment, &resul.Brand, &resul.Model, &resul.Sn)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}
func ReadDevice(id int64) (result []types.Device, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы устройств", err)
	}
	defer db.Close()
	res, err := db.Query("SELECT id, type, brand, model, sn from device WHERE id = ? ", id)
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице устройств")
	}
	for res.Next() {
		var resul types.Device
		err = res.Scan(&resul.Id, &resul.TypeEquipment, &resul.Brand, &resul.Model, &resul.Sn)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func NewDevice(device types.Device) (id int64, err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы мастеров", err)
	}
	defer db.Close()
	res, err := db.Exec("INSERT INTO `device` (type, brand, model, sn) VALUE (?, ?, ?, ?)", device.TypeEquipment, device.Brand, device.Model, device.Sn)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()

	return id, err

}
func ChangDevice(device types.Device) (err error) {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных с таблицы клиентов", err)
	}
	defer db.Close()
	_, err = db.Query("UPDATE `device` SET `type` = ?, `brand` = ?, `model` = ?, `sn` = ? WHERE `id` = ?", device.TypeEquipment, device.Brand, device.Model, device.Sn, device.Id)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return err
	}

	return err

}
func DelDevice(id int64) error {
	db, err := sql.Open("mysql", pass)
	if err != nil {
		fmt.Println("не удалось подключиться к базе данных для считывния данных для телеграм бота", err)
	}
	_, err = db.Query("DELETE FROM `device` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
