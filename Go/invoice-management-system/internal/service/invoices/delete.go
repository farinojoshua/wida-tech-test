package invoices

import (
	"context"
	"fmt"
)

func (s *service) DeleteInvoice(ctx context.Context, invoiceNo string) error {
	// check if invoice exists
	existingInvoice, err := s.invoiceRepo.GetByInvoiceNo(ctx, invoiceNo)
	if err != nil {
		return fmt.Errorf("failed to get invoice: %w", err)
	}
	if existingInvoice == nil {
		return fmt.Errorf("invoice not found")
	}

	// delete from database
	if err := s.invoiceRepo.Delete(ctx, invoiceNo); err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}
