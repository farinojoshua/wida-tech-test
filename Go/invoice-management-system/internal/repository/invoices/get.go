package invoices

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"invoice-management-system/internal/models"
)

func (r *repository) Get(ctx context.Context, startDate, endDate string, page, size int) ([]models.Invoice, error) {
	offset := (page - 1) * size

	query := `SELECT i.invoice_no, i.date, i.customer_name, i.salesperson_name, i.payment_type, i.notes, i.created_at, i.updated_at
  FROM invoices i
  WHERE DATE(i.date) BETWEEN ? AND ?
  ORDER BY i.created_at DESC
  LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, startDate, endDate, size, offset)
	if err != nil {
		return nil, fmt.Errorf("query invoices: %w", err)
	}
	defer rows.Close()

	var invoices []models.Invoice
	for rows.Next() {
		var invoice models.Invoice
		err := rows.Scan(
			&invoice.InvoiceNo,
			&invoice.Date,
			&invoice.CustomerName,
			&invoice.SalespersonName,
			&invoice.PaymentType,
			&invoice.Notes,
			&invoice.CreatedAt,
			&invoice.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan invoices: %w", err)
		}

		// Get products for each invoice
		products, err := r.getProducts(ctx, invoice.InvoiceNo)
		if err != nil {
			return nil, fmt.Errorf("failed to get products: %w", err)
		}

		invoice.Products = products

		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

func (r *repository) GetByInvoiceNo(ctx context.Context, invoiceNo string) (*models.Invoice, error) {
	query := `
    SELECT invoice_no, date, customer_name, salesperson_name, payment_type, notes, created_at, updated_at
    FROM invoices WHERE invoice_no = ?
  `

	var invoice models.Invoice
	err := r.db.QueryRowContext(ctx, query, invoiceNo).Scan(
		&invoice.InvoiceNo,
		&invoice.Date,
		&invoice.CustomerName,
		&invoice.SalespersonName,
		&invoice.PaymentType,
		&invoice.Notes,
		&invoice.CreatedAt,
		&invoice.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to get invoice: %w", err)
	}

	// get products
	products, err := r.getProducts(ctx, invoiceNo)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	invoice.Products = products

	return &invoice, nil
}

func (r *repository) GetMetadata(ctx context.Context, startDate, endDate string) (*models.Metadata, error) {
	query := `
  SELECT COUNT(*) as total_records,
      COUNT(CASE WHEN payment_type = 'CASH' THEN 1 END) as total_cash_transactions,
      COALESCE(SUM(
          (SELECT SUM(total_price_sold - total_cost_of_goods_sold)
          FROM products
          WHERE products.invoice_no = invoices.invoice_no)
      ), 0) as total_profit
  FROM invoices
  WHERE DATE(date) BETWEEN ? AND ?`

	var metadata models.Metadata
	err := r.db.QueryRowContext(ctx, query, startDate, endDate).Scan(
		&metadata.TotalRecords,
		&metadata.TotalCashTransactions,
		&metadata.TotalProfit,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	return &metadata, nil
}

// helper function to get product for an invoice
func (r *repository) getProducts(ctx context.Context, invoiceNo string) ([]models.Product, error) {
	query := `
    SELECT id, invoice_no, item_name, quantity, total_cost_of_goods_sold, total_price_sold, created_at, updated_at
    FROM products WHERE invoice_no = ?
  `

	rows, err := r.db.QueryContext(ctx, query, invoiceNo)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID,
			&product.InvoiceNo,
			&product.ItemName,
			&product.Quantity,
			&product.TotalCostOfGoodsSold,
			&product.TotalPriceSold,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}
