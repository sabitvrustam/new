package orders

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/sabitvrustam/new/pkg/types"
)

func (d *Order) ReadOrderParts(id *int64, idOrder int64) (orderP []types.OrderParts, err error) {
	sb := sq.Select("op.id", "op.id_orders", "op.id_parts", "p.parts_name", "p.parts_price").
		From("orders_parts AS op").
		Join("parts AS p ON op.id_parts = p.id")
	if id != nil {
		sb = sb.Where(sq.Eq{"op.id": *id})
	} else {
		sb = sb.Where(sq.Eq{"op.id_orders": idOrder})
	}
	res, err := sb.RunWith(d.db).Query()
	for res.Next() {
		var result types.OrderParts
		err = res.Scan(&result.Id, &result.IdOrder, &result.IdPart, &result.PartName, &result.PartPrice)
		if err != nil {
			d.log.Error(err)
		}
		orderP = append(orderP, result)
	}
	return orderP, err
}

func (d *Order) NewOrderParts(orderParts types.OrderParts) (id int64, err error) {
	sb := sq.Insert("orders_parts").
		Columns("id_orders", "id_parts").
		Values(orderParts.IdOrder, orderParts.IdPart)
	res, err := sb.RunWith(d.db).Exec()
	if err != nil {
		d.log.Error(err)
	}
	id, err = res.LastInsertId()
	if err != nil {
		d.log.Error(err)
	}
	return id, err
}

func (d *Order) DelOrderParts(id int64) (err error) {
	sb := sq.Delete("orders_parts").Where(sq.Eq{"id": id})
	_, err = sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err, "не удалось записать статус ")
	}
	return err
}
