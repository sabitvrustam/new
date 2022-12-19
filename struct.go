package main

type Order struct {
	User
	Device
	Works []Work
	Parts []Part
	Id
	Status
	Masters
	Price
	OllParts []Part
	OllWorks []Work
	Part
}

type User struct {
	FirstName string
	LastName  string
	MidlName  string
	Phone     string
}

type Device struct {
	TypeEquipment string
	Brand         string
	Model         string
	Sn            string
	Defect        string
}

type Id struct {
	IdOrder  string
	IdUser   string
	IdBrands string
}

type Work struct {
	Id        string
	IdWork    string
	WorkName  string
	WorkPrice string
}
type Part struct {
	Id         string
	IdPart     int
	PartsName  string
	PartsPrice string
}
type Status struct {
	StatusOrder string
}
type Price struct {
	PriceWork  string
	PriceParts string
}

type Masters struct {
	Id     string
	L_name string
}

type APIHandler struct {
	Id string
}
