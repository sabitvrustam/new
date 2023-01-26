package devices

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

func (d *Device) DevicesRead(id int64, sn string) (results []types.Device, err error) {
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

func (d *Device) DeviceWrite(device types.Device) (id int64, err error) {
	res := sq.Insert("device").
		Columns("type", "brand", "model", "sn").
		Values(device.TypeEquipment, device.Brand, device.Model, device.Sn)
	result, err := res.RunWith(d.db).Exec()
	if err != nil {
		fmt.Println("nooo")
	}
	id, err = result.LastInsertId()
	return id, err
}

func (d *Device) DeviceCange(device types.Device) (err error) {
	res := sq.Update("device").
		Set("type", device.TypeEquipment).
		Set("brand", device.Brand).
		Set("model", device.Model).
		Set("sn", device.Sn).
		Where(sq.Eq{"id": device.Id})
	_, err = res.RunWith(d.db).Exec()
	return err
}

func (d *Device) DeviceDelete(id int64) (err error) {
	res := sq.Delete("device").Where(sq.Eq{"id": id})
	_, err = res.RunWith(d.db).Exec()
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
