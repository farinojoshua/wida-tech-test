package invoices

import (
	"context"
	"database/sql"
	"fmt"
	"invoice-management-system/internal/models"
)

func (r *repository) Update(ctx context.Context, invoice *models.Invoice) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// update invoice
	query := `UPDATE invoices SET date = ?, customer_name = ?, salesperson_name = ?, payment_type = ?, notes = ? WHERE invoice_no = ?`

	result, err := tx.ExecContext(ctx, query,
		invoice.Date,
		invoice.CustomerName,
		invoice.SalespersonName,
		invoice.PaymentType,
		invoice.Notes,
		invoice.InvoiceNo,
	)

	if err != nil {
		return fmt.Errorf("failed to update invoice: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("invoice not found")
	}

	// Delete existing products
	_, err = tx.ExecContext(ctx, "DELETE FROM products WHERE invoice_no = ?", invoice.InvoiceNo)
	if err != nil {
		return fmt.Errorf("failed to delete existing products: %w", err)
	}

	// Insert new products
	for _, product := range invoice.Products {
		query = `INSERT INTO products (invoice_no, item_name, quantity, total_cost_of_goods_sold, total_price_sold) VALUES (?, ?, ?, ?, ?)`

		_, err = tx.ExecContext(ctx, query,
			invoice.InvoiceNo,
			product.ItemName,
			product.Quantity,
			product.TotalCostOfGoodsSold,
			product.TotalPriceSold,
		)
		if err != nil {
			return fmt.Errorf("failed to insert product: %w", err)
		}
	}

	return tx.Commit()
}
