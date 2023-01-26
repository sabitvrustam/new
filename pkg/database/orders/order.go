package orders

import (
	"database/sql"
	"fmt"

	//sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sabitvrustam/new/pkg/types"
	log "github.com/sirupsen/logrus"
)

type Order struct {
	db *sql.DB
}

func NewOrder(db *sql.DB) *Order {
	return &Order{db: db}
}

func (d *Order) ReadOrders(limit string, offset string) (Order []types.Order, err error) {
	var result types.Order
	res, err := d.db.Query("SELECT o.id, u.id, u.f_name, u.l_name, u.m_name, "+
		"u.n_phone, d.id, d.type, d.brand, d.model, d.sn, m.id, m.f_name, m.l_name, "+
		"m.m_name, m.n_phone, s.o_status FROM orders AS o "+
		"JOIN users AS u ON o.id_users = u.id "+
		"JOIN device AS d ON o.id_device = d.id "+
		"JOIN masters AS m ON o.id_masters  = m.id "+
		"JOIN status AS s ON o.id_status  = s.id "+
		"ORDER BY o.id DESC  LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		log.Error("не удалось считать данные заказа из базы данных", err)
	}
	log.Info("hello")
	for res.Next() {
		err = res.Scan(&result.IdOrder, &result.User.Id, &result.User.FirstName, &result.User.LastName,
			&result.User.MidlName, &result.User.Phone, &result.Device.Id, &result.TypeEquipment,
			&result.Brand, &result.Model, &result.Sn, &result.Master.Id, &result.Master.FirstName,
			&result.Master.LastName, &result.Master.MidlName, &result.Master.Phone, &result.Status.StatusOrder)
		if err != nil {
			fmt.Println(err)
		}
		Order = append(Order, result)
	}
	return Order, err

}

func (d *Order) ReadOrder(id int64) (Order types.Order, err error) {
	var result types.Order
	res, err := d.db.Query("SELECT o.id, u.id, u.f_name, u.l_name, u.m_name, "+
		"u.n_phone, d.id, d.type, d.brand, d.model, d.sn, m.id, m.f_name, m.l_name, "+
		"m.m_name, m.n_phone, s.o_status FROM orders AS o "+
		"JOIN users AS u ON o.id_users = u.id "+
		"JOIN device AS d ON o.id_device = d.id "+
		"JOIN masters AS m ON o.id_masters  = m.id "+
		"JOIN status AS s ON o.id_status  = s.id "+
		"WHERE o.id = ?", id)
	if err != nil {
		fmt.Sprintln("не удалось считать данные заказа из базы данных", err)
	}
	for res.Next() {
		err = res.Scan(&result.IdOrder, &result.User.Id, &result.User.FirstName, &result.User.LastName,
			&result.User.MidlName, &result.User.Phone, &result.Device.Id, &result.TypeEquipment,
			&result.Brand, &result.Model, &result.Sn, &result.Master.Id, &result.Master.FirstName,
			&result.Master.LastName, &result.Master.MidlName, &result.Master.Phone, &result.Status.StatusOrder)
		if err != nil {
			fmt.Println(err)
		}
	}
	return result, err
}
func (d *Order) NewOrder1(order types.Id) (id int64, err error) {
	res, err := d.db.Exec("INSERT INTO `orders` (`id_users`, `id_device`, `id_masters`, `id_status`) VALUE (?, ?, ?, ?)", order.IdUser, order.IdDevice, order.IdMaster, order.IdStatus)
	if err != nil {
		fmt.Println("не удалось записать ключи в таблицу заказов", err)
		return 0, err
	}
	id, err = res.LastInsertId()
	return id, err
}
func (d *Order) ChangeOrder(order types.Id) (err error) {
	_, err = d.db.Exec("UPDATE `orders`"+
		"SET `id_users` = ?, `id_device` = ?, `id_masters` = ?, `id_status` = ?"+
		" WHERE `id` = ?", order.IdUser, order.IdDevice, order.IdMaster, order.IdStatus, order.IdOrder)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
	}
	return
}
func (d *Order) DelOrder(id int64) (err error) {
	_, err = d.db.Query("DELETE FROM `orders` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return
}
