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
	IdOrder  string `json:"id_order"`
	IdUser   string `json:"id_user"`
	IdBrands string `json:"id_brands"`
}

type Work struct {
	Id        string `json:"id"`
	IdWork    string `json:"id_work"`
	WorkName  string `json:"work_name"`
	WorkPrice string `json:"work_price"`
}
type Part struct {
	Id         string `json:"id"`
	IdPart     int    `json:"id_part"`
	PartsName  string `json:"parts_name"`
	PartsPrice string `json:"parts_price"`
}
type Status struct {
	StatusOrder string `json:"status_order"`
}
type Price struct {
	PriceWork  string `json:"price_work"`
	PriceParts string `json:"price_parts"`
}

type Masters struct {
	Id     string `json:"id"`
	L_name string `json:"l_name"`
}

type APIHandler struct {
	Id string
}
