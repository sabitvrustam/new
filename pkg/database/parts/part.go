package parts

import (
	"database/sql"
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

type Part struct {
	db *sql.DB
}

func NewPart(db *sql.DB) *Part {
	return &Part{db: db}
}

func (d *Part) ReadParts() (result []types.Part) {
	res, err := d.db.Query("SELECT id, parts_name, parts_price from parts ")
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {
		var resul types.Part
		err = res.Scan(&resul.Id, &resul.PartsName, &resul.PartsPrice)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}

	return result
}

func (d *Part) NewPart1(newPart types.Part) (id int64, err error) {
	res, err := d.db.Exec("INSERT INTO `parts` (`parts_name`, `parts_price`) VALUE (?, ?)", newPart.PartsName, newPart.PartsPrice)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()

	return id, err
}

func (d *Part) ChangePart(part types.Part) (err error) {
	_, err = d.db.Query("UPDATE `parts` SET `parts_name` = ?, `parts_price` = ?  WHERE `id` = ?", part.PartsName, part.PartsPrice, part.Id)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return err
	}

	return err

}

func (d *Part) DeletePart(id int64) (err error) {
	_, err = d.db.Query("DELETE FROM `parts` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err

}
