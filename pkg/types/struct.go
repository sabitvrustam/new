package types

type Order struct {
	User    `json:"user"`
	Device  `json:"device"`
	Works   []Work
	Parts   []Part
	Id      `json:"id"`
	Status  `json:"status"`
	Masters `json:"masters"`
	Price   `json:"price"`
}

type User struct {
	Id        int64  `json:"id_user"`
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
	IdOrder  int64 `json:"id_order"`
	IdUser   int64 `json:"id_user"`
	IdBrands int64 `json:"id_brands"`
}

type Work struct {
	Id        int64  `json:"id"`
	WorkName  string `json:"work_name"`
	WorkPrice string `json:"work_price"`
}
type Part struct {
	Id         int64  `json:"id"`
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
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	MidlName  string `json:"midl_name"`
	Phone     string `json:"phone"`
}
