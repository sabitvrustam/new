package main

type Order struct {
	User     `json:"user"`
	Device   `json:"device"`
	Works    []Work
	Parts    []Part
	Id       `json:"id"`
	Status   `json:"status"`
	Masters  `json:"masters"`
	Price    `json:"price"`
	OllParts []Part
	OllWorks []Work
	Part     `json:"part"`
	Work     `json:"work"`
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	MidlName  string `json:"midl_name"`
	Phone     string `json:"phone"`
}

type Device struct {
	TypeEquipment string `json:"type_equipment"`
	Brand         string `json:"brand"`
	Model         string `json:"model"`
	Sn            string `json:"sn"`
	Defect        string `json:"defect"`
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
