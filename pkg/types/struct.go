package types

type Order struct {
	IdOrder int64 `json:"id_order"`
	User    `json:"user"`
	Device  `json:"device"`
	Status  `json:"status"`
	Master  `json:"master"`
}
type User struct {
	Id        int64  `json:"id_user"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	MidlName  string `json:"midl_name"`
	Phone     string `json:"phone"`
}
type Device struct {
	Id            int64  `json:"id_device"`
	TypeEquipment string `json:"type_equipment"`
	Brand         string `json:"brand"`
	Model         string `json:"model"`
	Sn            string `json:"sn"`
	Defect        string `json:"defect"`
}

type Id struct {
	IdOrder  int64 `json:"id_order"`
	IdUser   int64 `json:"id_user"`
	IdDevice int64 `json:"id_device"`
	IdMaster int64 `json:"id_master"`
	IdStatus int64 `json:"id_status"`
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
	Id          int64  `json:"id_status"`
	StatusOrder string `json:"status_order"`
}
type Price struct {
	PriceWork  string `json:"price_work"`
	PriceParts string `json:"price_parts"`
}

type Master struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	MidlName  string `json:"midl_name"`
	Phone     string `json:"phone"`
}
type OrderParts struct {
	Id        int64  `json:"id"`
	IdOrder   int64  `json:"id_order"`
	IdPart    int64  `json:"id_part"`
	PartName  string `json:"part_name"`
	PartPrice string `json:"part_price"`
}
type OrderWorks struct {
	Id        int64  `json:"id"`
	IdOrder   int64  `json:"id_order"`
	IdWork    int64  `json:"id_work"`
	WorkName  string `json:"work_name"`
	WorkPrice string `json:"work_price"`
}
