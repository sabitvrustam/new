package main

import "database/sql"

type UserRead struct {
	FirstName string
	LastName  string
	Phone     string
}

type UserWrite struct {
	firstName string
	lastName  string
	phone     string
}

type Equipment struct {
	typeEquipment string
	brand         string
	model         string
	sn            string
}

type DataWrite struct {
	db *sql.DB
}

type DataRead struct {
	db *sql.DB
}
