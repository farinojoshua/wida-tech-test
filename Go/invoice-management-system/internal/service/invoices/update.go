package invoices

import (
	"context"
	"fmt"
	"invoice-management-system/internal/models"
	"time"
)

func (s *service) UpdateInvoice(ctx context.Context, invoiceNo string, req *models.UpdateInvoiceRequest) error {
	// Validate request
	if errs := req.Validate(); len(errs) > 0 {
		return fmt.Errorf("validation error: %v", errs)
	}

	// Check if invoice exists
	existingInvoice, err := s.invoiceRepo.GetByInvoiceNo(ctx, invoiceNo)
	if err != nil {
		return fmt.Errorf("failed to get invoice: %w", err)
	}
	if existingInvoice == nil {
		return fmt.Errorf("invoice not found")
	}

	// Parse date
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return fmt.Errorf("invalid date format: %w", err)
	}

	// Create invoice model
	invoice := &models.Invoice{
		InvoiceNo:       invoiceNo,
		Date:            date,
		CustomerName:    req.CustomerName,
		SalespersonName: req.SalespersonName,
		PaymentType:     req.PaymentType,
		Notes:           req.Notes,
	}

	// Map products
	for _, p := range req.Products {
		invoice.Products = append(invoice.Products, models.Product{
			InvoiceNo:            invoiceNo,
			ItemName:             p.ItemName,
			Quantity:             p.Quantity,
			TotalCostOfGoodsSold: p.TotalCostOfGoodsSold,
			TotalPriceSold:       p.TotalPriceSold,
		})
	}

	// Update in database
	if err := s.invoiceRepo.Update(ctx, invoice); err != nil {
		return fmt.Errorf("failed to update invoice: %w", err)
	}

	return nil
}
