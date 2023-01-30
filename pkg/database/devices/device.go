package devices

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type Device struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewDevice(db *sql.DB, log *logrus.Logger) *Device {
	return &Device{
		db:  db,
		log: log}
}

func (d *Device) DevicesRead(id int64, sn string) (results []types.Device, err error) {
	sb := sq.Select(" id, type, brand, model, sn").From("device")
	if id != 0 && sn == "" {
		sb = sb.Where(sq.Eq{"id": id})
	}
	if id == 0 && sn != "" {
		sb = sb.Where(sq.Eq{"sn": sn})
	}
	res, err := sb.RunWith(d.db).Query()
	for res.Next() {
		var resul types.Device
		err = res.Scan(&resul.Id, &resul.TypeEquipment, &resul.Brand, &resul.Model, &resul.Sn)
		if err != nil {
			d.log.Error(err)
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
		d.log.Error(err, "nooo")
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
	if err != nil {
		d.log.Error(err)
	}
	return err
}

func (d *Device) DeviceDelete(id int64) (err error) {
	res := sq.Delete("device").Where(sq.Eq{"id": id})
	_, err = res.RunWith(d.db).Exec()
	if err != nil {
		d.log.Error(err, "не удалось записать статус ")
	}
	return err
}
