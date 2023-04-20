package models

type ProductWh struct {
	ID              int `json:"id"`
	ProductName     string `json:"product_name"`
	Price           int    `json:"price"`
	ProductCategory string `json:"product_category"`
	Stock           int    `json:"stock"`
}