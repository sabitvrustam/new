package orders

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/sabitvrustam/new/pkg/types"
)

func (d *Order) ReadOrderWorks(idOrder int64, id *int64) (result []types.OrderWorks, err error) {
	sb := sq.Select("ow.id", "ow.id_orders", "ow.id_work", "w.work_name", "w.work_price").
		From("orders_work AS ow").
		Join("work AS w ON ow.id_work = w.id")
	if id != nil {
		sb = sb.Where(sq.Eq{"ow.id": *id})
	} else {
		sb = sb.Where(sq.Eq{"ow.id_orders": idOrder})
	}
	res, err := sb.RunWith(d.db).Query()
	for res.Next() {
		var resul types.OrderWorks
		err = res.Scan(&resul.Id, &resul.IdOrder, &resul.IdWork, &resul.WorkName, &resul.WorkPrice)
		if err != nil {
			d.log.Error(err)
		}
		result = append(result, resul)
	}
	return result, err
}

func (d *Order) NewOrderWorks(orderWorks types.OrderWorks) (id int64, err error) {
	sb := sq.Insert("orders_work").
		Columns("id_orders", "id_work").
		Values(orderWorks.IdOrder, orderWorks.IdWork)
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

func (d *Order) DelOrderWorks(id int64) (err error) {
	sb := sq.Delete("orders_work").Where(sq.Eq{"id": id})
	_, err = sb.RunWith(d.db).Query()
	if err != nil {
		d.log.Error(err)
	}
	return err
}
