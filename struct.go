package main

import "database/sql"

type User struct {
	FirstName string
	LastName  string
	MidlName  string
	Phone     string
}

type Equipment struct {
	TypeEquipment string
	Brand         string
	Model         string
	Sn            string
}

type Id struct {
	IdOrder  string
	IdUser   string
	IdBrands string
}

type DataWrite struct {
	db *sql.DB
}

type DataRead struct {
	db *sql.DB
}
