package models

type CreateInvoiceRequest struct {
	InvoiceNo       string      `json:"invoice_no" binding:"required"`
	Date            string      `json:"date" binding:"required"`
	CustomerName    string      `json:"customer_name" binding:"required,min=2"`
	SalespersonName string      `json:"salesperson_name" binding:"required,min=2"`
	PaymentType     PaymentType `json:"payment_type" binding:"required,oneof=CASH CREDIT"`
	Notes           *string     `json:"notes,omitempty" binding:"omitempty,min=5"`
	Products        []struct {
		ItemName             string  `json:"item_name" binding:"required,min=5"`
		Quantity             int     `json:"quantity" binding:"required,min=1"`
		TotalCostOfGoodsSold float64 `json:"total_cost_of_goods_sold" binding:"required,min=0"`
		TotalPriceSold       float64 `json:"total_price_sold" binding:"required,min=0"`
	} `json:"products" binding:"required,dive"`
}

type UpdateInvoiceRequest CreateInvoiceRequest

type GetInvoicesRequest struct {
	StartDate string `form:"start_date" binding:"required"`
	EndDate   string `form:"end_date" binding:"required"`
	Page      int    `form:"page" binding:"required,min=1"`
	Size      int    `form:"size" binding:"required,min=1"`
}
