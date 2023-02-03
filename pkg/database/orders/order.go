package orders

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type Order struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewOrder(db *sql.DB, log *logrus.Logger) *Order {
	return &Order{
		db:  db,
		log: log}
}

func (d *Order) GetOrderByID(id int64) (Order *types.Order, err error) {
	orders, err := d.readOrders(0, 0, &id)
	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}
	return orders[0], err
}

func (d *Order) ReadOrders(limit uint64, offset uint64) (orders []*types.Order, err error) {
	return d.readOrders(limit, offset, nil)
}

func (d *Order) readOrders(limit uint64, offset uint64, id *int64) (orders []*types.Order, err error) {
	sb := sq.Select("o.id", "u.id", "u.f_name", "u.l_name", "u.m_name", "u.n_phone",
		"d.id", "d.type", "d.brand", "d.model", "d.sn",
		"m.id", "m.f_name", "m.l_name", "m.m_name", "m.n_phone", "s.id", "s.o_status").
		From("orders AS o").
		Join("users AS u ON o.id_users = u.id").
		Join("device AS d ON o.id_device = d.id").
		Join("masters AS m ON o.id_masters  = m.id").
		Join("status AS s ON o.id_status  = s.id")
	if id != nil {
		sb = sb.Where(sq.Eq{"o.id": *id})
	} else {
		sb = sb.OrderBy("o.id desc").Limit(limit).Offset(offset)
	}
	res, err := sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err, "не удалось считать заказы из базы данных")
	}
	for res.Next() {
		var result types.Order
		err = res.Scan(&result.IdOrder, &result.User.Id, &result.User.FirstName, &result.User.LastName,
			&result.User.MidlName, &result.User.Phone, &result.Device.Id, &result.TypeEquipment,
			&result.Brand, &result.Model, &result.Sn, &result.Master.Id, &result.Master.FirstName,
			&result.Master.LastName, &result.Master.MidlName, &result.Master.Phone, &result.Status.Id, &result.Status.StatusOrder)
		if err != nil {
			d.log.Error(err)
		}
		orders = append(orders, &result)
	}
	return orders, err
}

func (d *Order) NewOrder1(order types.Id) (id int64, err error) {
	orders := sq.Insert("orders").
		Columns("id_users", "id_device", "id_masters", "id_status").
		Values(order.IdUser, order.IdDevice, order.IdMaster, order.IdStatus)
	res, err := orders.RunWith(d.db).Exec()
	if err != nil {
		d.log.Error("не удалось записать данные заказа в базу данных", err)
	}
	id, err = res.LastInsertId()
	return id, err
}

func (d *Order) ChangeOrder(order types.Id) (err error) {
	orders := sq.Update("orders").
		Set("id_users", order.IdUser).
		Set("id_device", order.IdDevice).
		Set("id_masters", order.IdMaster).
		Set("id_status", order.IdStatus).
		Where(sq.Eq{"id": order.IdOrder})
	_, err = orders.RunWith(d.db).Exec()
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
	}
	return
}

func (d *Order) DelOrder(id int64) (err error) {
	_, err = sq.Delete("orders").Where(sq.Eq{"o.id": id}).RunWith(d.db).Exec()
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return
}
