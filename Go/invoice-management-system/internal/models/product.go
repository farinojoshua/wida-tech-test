package models

import (
	"time"
)

type Product struct {
	ID                   int64     `json:"id" db:"id"`
	InvoiceNo            string    `json:"invoice_no" db:"invoice_no"`
	ItemName             string    `json:"item_name" db:"item_name"`
	Quantity             int       `json:"quantity" db:"quantity"`
	TotalCostOfGoodsSold float64   `json:"total_cost_of_goods_sold" db:"total_cost_of_goods_sold"`
	TotalPriceSold       float64   `json:"total_price_sold" db:"total_price_sold"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
}
