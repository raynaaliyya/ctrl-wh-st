package entity

type ReportWh struct {
	ProductWhId string `json:"product_wh_id"`
	Stock       int    `json:"stock"`
	ProductName string `json:"product_name"`
	Act         string `json:"act"`
	LastStock   int    `json:"last_stock"`
	CreatedAt   string `json:"created_at"`
}
