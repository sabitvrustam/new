package works

import (
	"database/sql"
	"fmt"

	"github.com/sabitvrustam/new/pkg/types"
)

type Work struct {
	db *sql.DB
}

func NewWork(db *sql.DB) *Work {
	return &Work{db: db}
}

func (d *Work) ReadWoeks() (result []types.Work) {
	res, err := d.db.Query("SELECT id, work_name, work_price from work ")
	if err != nil {
		fmt.Sprintln(err)
	}
	for res.Next() {
		var resul types.Work
		err = res.Scan(&resul.Id, &resul.WorkName, &resul.WorkPrice)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, resul)
	}

	return result
}

func (d *Work) WriteWork(newWork types.Work) (id int64, err error) {
	res, err := d.db.Exec("INSERT INTO `work` (`work_name`, `work_price`) VALUE (?, ?)", newWork.WorkName, newWork.WorkPrice)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
		return 0, err
	}
	id, err = res.LastInsertId()
	return id, err
}

func (d *Work) ChangeWork(work types.Work) (err error) {
	_, err = d.db.Query("UPDATE `work` SET `work_name` = ?, `work_price` = ?  WHERE `id` = ?", work.WorkName, work.WorkPrice, work.Id)
	if err != nil {
		fmt.Println("не удалось записать новую запчасть в базу данных", err)
	}
	return err
}

func (d *Work) DeleteWork(id int64) (err error) {
	_, err = d.db.Query("DELETE FROM `work` WHERE `id`=?", id)
	if err != nil {
		fmt.Println(err, "не удалось записать статус ")
	}
	return err

}
