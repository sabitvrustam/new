package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func Migrate(db *sql.DB, log *logrus.Logger) {

	var dbuser string = os.Getenv("bduser")
	var dbpass string = os.Getenv("bdpass")
	var pass string = fmt.Sprintf("%s:%s@tcp(185.68.93.47:3306)/", dbuser, dbpass)
	d, err := sql.Open("mysql", pass)
	if err != nil {
		log.Error("не удалось подключиться к базе данных", err)
	} else {
		log.Info("Подключение к базе данных")
	}

	res, err := d.Query("show databases")
	if err != nil {
		fmt.Println(err, "не удалось выполнить команду mysql показать базы данных")
	}
	var dbStatus bool
	for res.Next() {
		var result string
		err := res.Scan(&result)
		if err != nil {
			fmt.Println(err, "не удалось записать данные в переменную")
		}
		if result == "my_service" {
			dbStatus = true
		}
		fmt.Println(result)
	}

	if !dbStatus {
		_, err := d.Query("CREATE DATABASE `my_service`")
		if err != nil {
			fmt.Println(err, "не удалось выполнить команду mysql создания базы данных")
			return
		}
		_, err = db.Query("CREATE TABLE `users` ( " +
			"`id` int NOT NULL AUTO_INCREMENT, " +
			"`l_name` varchar(20) NOT NULL, " +
			"`f_name` varchar(20) NOT NULL, " +
			"`m_name` varchar(20) NOT NULL, " +
			"`n_phone` varchar(12) NOT NULL, " +
			"`id_telegram` varchar(20) DEFAULT NULL, " +
			"PRIMARY KEY (`id`) " +
			")")
		if err != nil {
			fmt.Println(res, err, "не удалось выполнить команду mysql создания таблицы клиентов")
		} else {
			fmt.Println("таблица клиентов создана")
		}
		_, err = db.Query("CREATE TABLE `device` ( " +
			"`id` int NOT NULL AUTO_INCREMENT, " +
			"`type` varchar(20) NOT NULL, " +
			"`brand` varchar(20) NOT NULL, " +
			"`model` varchar(20) NOT NULL, " +
			"`sn` varchar(20) NOT NULL, " +
			"PRIMARY KEY (`id`) " +
			")")
		if err != nil {
			fmt.Println(res, err, "не удалось выполнить команду mysql создания таблицы устроиств")
		} else {
			fmt.Println("таблица устроиств создана")
		}
		_, err = db.Query("CREATE TABLE `status` ( " +
			"`id` int NOT NULL AUTO_INCREMENT, " +
			"`o_status` varchar(20) NOT NULL, " +
			"PRIMARY KEY (`id`) " +
			")")
		if err != nil {
			fmt.Println(res, err, "не удалось выполнить команду mysql создания таблицы соятояния")
		} else {
			fmt.Println("таблица состояния заказа создана")
		}
		status := []string{"Принято", "На диагностике", "На соглосовании", "Отказ от ремонта", "Ждет запчасть", "В работе",
			"Готов к выдаче", "Выдано"}
		kol := len(status)
		for i, statusName := range status {
			_, err = db.Query("INSERT INTO `status` (`o_status`) VALUE (?)", statusName)
			if err != nil {
				fmt.Println(res, err, "не удалось записать статус ", i)
			} else {
				fmt.Println("Данные в таблицу статусов записаны, в колличестве", kol)
			}
		}
		_, err = db.Query("CREATE TABLE `masters` ( " +
			"`id` int NOT NULL AUTO_INCREMENT, " +
			"`l_name` varchar(20) NOT NULL, " +
			"`f_name` varchar(20) NOT NULL, " +
			"`m_name` varchar(20) NOT NULL, " +
			"`n_phone` varchar(12) NOT NULL, " +
			"`id_telegram` varchar(20) DEFAULT NULL, " +
			"PRIMARY KEY (`id`) " +
			")")
		if err != nil {
			fmt.Println(res, err, "не удалось выполнить команду mysql создания таблицы мастеров")
		} else {
			fmt.Println("таблица мастеров создана")
		}
		res, err = db.Query("CREATE TABLE `orders` ( " +
			"`id` int NOT NULL AUTO_INCREMENT, " +
			"`id_users` int NOT NULL, " +
			"`id_device` int NOT NULL, " +
			"`id_status` int DEFAULT NULL, " +
			"`id_masters` int DEFAULT NULL, " +
			"PRIMARY KEY (`id`), " +
			"KEY `orders_FK` (`id_users`), " +
			"KEY `orders_FK_1` (`id_device`), " +
			"KEY `orders_FK_2` (`id_status`), " +
			"KEY `orders_FK_3` (`id_masters`), " +
			"CONSTRAINT `orders_FK` FOREIGN KEY (`id_users`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT, " +
			"CONSTRAINT `orders_FK_1` FOREIGN KEY (`id_device`) REFERENCES `device` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT, " +
			"CONSTRAINT `orders_FK_2` FOREIGN KEY (`id_status`) REFERENCES `status` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT, " +
			"CONSTRAINT `orders_FK_3` FOREIGN KEY (`id_masters`) REFERENCES `masters` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT " +
			")")
		if err != nil {
			fmt.Println(res, err, "не удалось выполнить команду mysql создания таблицы заказов")
		} else {
			fmt.Println("таблица заказов создана")
		}

		_, err = db.Query("CREATE TABLE `parts` ( " +
			"`id` int NOT NULL AUTO_INCREMENT, " +
			"`parts_name` varchar(20) NOT NULL, " +
			"`parts_price` varchar(20) NOT NULL, " +
			"PRIMARY KEY (`id`) " +
			")")
		if err != nil {
			fmt.Println(res, err, "не удалось выполнить команду mysql создания таблицы запчастей")
		} else {
			fmt.Println("таблица запчастей создана")
		}
		_, err = db.Query("CREATE TABLE `work` ( " +
			"`id` int NOT NULL AUTO_INCREMENT, " +
			"`work_name` varchar(20) NOT NULL, " +
			"`work_price` varchar(20) NOT NULL, " +
			"PRIMARY KEY (`id`) " +
			")")
		if err != nil {
			fmt.Println(res, err, "не удалось выполнить команду mysql создания таблицы работ")
		} else {
			fmt.Println("таблица работ создана")
		}
		_, err = db.Query("CREATE TABLE `orders_parts` ( " +
			"`id` int NOT NULL AUTO_INCREMENT, " +
			"`id_orders` int NOT NULL, " +
			"`id_parts` int NOT NULL, " +
			"PRIMARY KEY (`id`), " +
			"KEY `orders_parts_FK` (`id_orders`), " +
			"KEY `orders_parts_FK_1` (`id_parts`), " +
			"CONSTRAINT `orders_parts_FK` FOREIGN KEY (`id_orders`) REFERENCES `orders` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT, " +
			"CONSTRAINT `orders_parts_FK_1` FOREIGN KEY (`id_parts`) REFERENCES `parts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT " +
			")")
		if err != nil {
			fmt.Println(res, err, "не удалось выполнить команду mysql создания таблицы заказ_запчасти")
		} else {
			fmt.Println("таблица заказ_запчасти создана")
		}

		_, err = db.Query("CREATE TABLE `orders_work` ( " +
			"`id` int NOT NULL AUTO_INCREMENT, " +
			"`id_orders` int NOT NULL, " +
			"`id_work` int NOT NULL, " +
			"PRIMARY KEY (`id`), " +
			"KEY `orders_work_FK` (`id_orders`), " +
			"KEY `orders_work_FK_1` (`id_work`), " +
			"CONSTRAINT `orders_work_FK` FOREIGN KEY (`id_orders`) REFERENCES `orders` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT, " +
			"CONSTRAINT `orders_work_FK_1` FOREIGN KEY (`id_work`) REFERENCES `work` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT " +
			")")
		if err != nil {
			fmt.Println("не удалось выполнить команду mysql создания таблицы заказ_работы", res, err)
		} else {
			fmt.Println("таблица заказ_работы создана")
		}

	} else {
		log.Info("база данных существует")
	}
}
