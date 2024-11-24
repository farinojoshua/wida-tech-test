package invoices

import (
	"context"
	"fmt"
)

func (r *repository) Delete(ctx context.Context, invoiceNo string) error {
	// Products will be deleted automatically when the invoice is deleted
	query := "DELETE FROM invoices WHERE invoice_no = ?"

	result, err := r.db.ExecContext(ctx, query, invoiceNo)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("invoice not found")
	}

	return nil
}
