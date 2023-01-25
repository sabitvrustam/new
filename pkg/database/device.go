package database

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/sabitvrustam/new/pkg/types"
)

type Device struct {
	db *sql.DB
}

func NewDevice(db *sql.DB) *Device {
	return &Device{db: db}
}

func (d *Device) ReadDevices(id int64, sn string) (results []types.Device, err error) {
	var res *sql.Rows
	devices := sq.Select(" id, type, brand, model, sn").From("device")
	activeId := devices.Where(sq.Eq{"id": id})
	activeSN := devices.Where(sq.Eq{"sn": sn})
	if id == 0 && sn == "" {
		res, err = devices.RunWith(d.db).Query()
	}
	if id != 0 && sn == "" {
		res, err = activeId.RunWith(d.db).Query()
	}
	if id == 0 && sn != "" {
		res, err = activeSN.RunWith(d.db).Query()
	}
	if err != nil {
		fmt.Println("upsss", err)
		return
	}
	for res.Next() {
		var resul types.Device
		err = res.Scan(&resul.Id, &resul.TypeEquipment, &resul.Brand, &resul.Model, &resul.Sn)
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, resul)
	}
	return results, err
}

func (d *Device) NewDevice1(device types.Device) (id int64, err error) {
	res, err := d.db.Exec("INSERT INTO `device` (type, brand, model, sn) VALUE (?, ?, ?, ?)", device.TypeEquipment, device.Brand, device.Model, device.Sn)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()
	return id, err
}

func (d *Device) ChangDevice(device types.Device) (err error) {
	_, err = d.db.Query("UPDATE `device` SET `type` = ?, `brand` = ?, `model` = ?, `sn` = ? WHERE `id` = ?", device.TypeEquipment, device.Brand, device.Model, device.Sn, device.Id)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return err
	}
	return err
}

func (d *Device) DelDevice(id int64) error {
	_, err := d.db.Query("DELETE FROM `device` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
