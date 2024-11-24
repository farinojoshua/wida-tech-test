package models

import (
	"time"
)

type PaymentType string

const (
	PaymentTypeCash   PaymentType = "CASH"
	PaymentTypeCredit PaymentType = "CREDIT"
)

type Invoice struct {
	InvoiceNo       string      `json:"invoice_no" db:"invoice_no"`
	Date            time.Time   `json:"date" db:"date"`
	CustomerName    string      `json:"customer_name" db:"customer_name"`
	SalespersonName string      `json:"salesperson_name" db:"salesperson_name"`
	PaymentType     PaymentType `json:"payment_type" db:"payment_type"`
	Notes           *string     `json:"notes,omitempty" db:"notes"`
	Products        []Product   `json:"products"`
	CreatedAt       time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at" db:"updated_at"`
}
