package models

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []Error     `json:"errors,omitempty"`
}

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type GetInvoicesResponse struct {
	Invoices []Invoice `json:"invoices"`
	Metadata Metadata  `json:"metadata"`
}

type Metadata struct {
	TotalProfit           float64 `json:"total_profit"`
	TotalCashTransactions int     `json:"total_cash_transactions"`
	TotalRecords          int     `json:"total_records"`
	TotalPages            int     `json:"total_pages"`
	CurrentPage           int     `json:"current_page"`
}

type CreateInvoiceResponse struct {
	InvoiceNo string `json:"invoice_no"`
}

type ImportInvoiceResponse struct {
	ImportedCount int           `json:"imported_count"`
	Errors        []ImportError `json:"errors,omitempty"`
}

type ImportError struct {
	InvoiceNo string `json:"invoice_no"`
	Message   string `json:"message"`
}
