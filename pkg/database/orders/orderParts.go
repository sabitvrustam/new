package orders

import (
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

func (d *Order) ReadOrderPart(id int64) (result []types.OrderParts, err error) {
	res, err := d.db.Query("SELECT op.id, op.id_orders, op.id_parts, p.parts_name, p.parts_price from orders_parts AS op "+
		"JOIN parts AS p ON op.id_parts = p.id "+
		"WHERE op.id =?", id)
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице устройств")
	}
	for res.Next() {
		var resul types.OrderParts
		err = res.Scan(&resul.Id, &resul.IdOrder, &resul.IdPart, &resul.PartName, &resul.PartPrice)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func (d *Order) ReadOrderParts(idOrder int64) (result []types.OrderParts, err error) {
	res, err := d.db.Query("SELECT op.id, op.id_orders, op.id_parts, p.parts_name, p.parts_price from orders_parts AS op "+
		"JOIN parts AS p ON op.id_parts = p.id "+
		"WHERE op.id_orders =?", idOrder)
	if err != nil {
		fmt.Sprintln(err, "не удалось выполнить запрос селект к таблице устройств")
	}
	for res.Next() {
		var resul types.OrderParts
		err = res.Scan(&resul.Id, &resul.IdOrder, &resul.IdPart, &resul.PartName, &resul.PartPrice)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func (d *Order) NewOrderParts(orderParts types.OrderParts) (id int64, err error) {
	res, err := d.db.Exec("INSERT INTO `orders_parts` (id_orders, id_parts) VALUE (?, ?)", orderParts.IdOrder,
		orderParts.IdPart)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()

	return id, err

}
func (d *Order) DelOrderParts(id int64) (err error) {
	_, err = d.db.Query("DELETE FROM `Orders_parts` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err
}
