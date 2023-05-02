package models

import "time"

type TrxWh struct {
	ID          string    `json:"id"`
	ProductWhId string    `json:"product_wh_id"`
	ProductName string    `json:"product_name"`
	Stock       int       `json:"stock"`
	Act         string    `json:"act"`
	CreatedAt   time.Time `json:"created_at"`
}
