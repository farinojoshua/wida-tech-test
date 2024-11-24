package invoices

import (
	"context"
	"database/sql"
	"fmt"
	"invoice-management-system/internal/models"
)

func (r *repository) Create(ctx context.Context, invoice *models.Invoice) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// insert invoice
	query := `INSERT INTO invoices (invoice_no, date, customer_name, salesperson_name, payment_type, notes) VALUES (?, ?, ?, ?, ?, ?)`

	_, err = tx.ExecContext(ctx, query, invoice.InvoiceNo, invoice.Date, invoice.CustomerName, invoice.SalespersonName, invoice.PaymentType, invoice.Notes)
	if err != nil {
		return fmt.Errorf("failed to insert invoice: %w", err)
	}

	// insert products
	for _, product := range invoice.Products {
		query = `INSERT INTO products (invoice_no, item_name, quantity, total_cost_of_goods_sold, total_price_sold) VALUES (?, ?, ?, ?, ?)`

		_, err = tx.ExecContext(ctx, query, invoice.InvoiceNo, product.ItemName, product.Quantity, product.TotalCostOfGoodsSold, product.TotalPriceSold)
		if err != nil {
			return fmt.Errorf("failed to insert product: %w", err)
		}
	}

	return tx.Commit()
}
