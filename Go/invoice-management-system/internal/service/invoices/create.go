package invoices

import (
	"context"
	"fmt"
	"invoice-management-system/internal/models"
	"time"
)

func (s *service) CreateInvoice(ctx context.Context, req *models.CreateInvoiceRequest) (*models.CreateInvoiceResponse, error) {
	// validate request
	if errs := req.Validate(); len(errs) > 0 {
		return nil, fmt.Errorf("validation error: %v", errs)
	}

	// parse date
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format, must be YYYY-MM-DD: %w", err)
	}

	// check if invoice number already exists
	existingInvoice, err := s.invoiceRepo.GetByInvoiceNo(ctx, req.InvoiceNo)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing invoice: %w", err)
	}
	if existingInvoice != nil {
		return nil, fmt.Errorf("invoice number %s already exists", req.InvoiceNo)
	}

	// create invoice
	invoice := &models.Invoice{
		InvoiceNo:       req.InvoiceNo,
		Date:            date,
		CustomerName:    req.CustomerName,
		SalespersonName: req.SalespersonName,
		PaymentType:     req.PaymentType,
		Notes:           req.Notes,
	}

	// map products
	for _, product := range req.Products {
		invoice.Products = append(invoice.Products, models.Product{
			ItemName:             product.ItemName,
			Quantity:             product.Quantity,
			TotalCostOfGoodsSold: product.TotalCostOfGoodsSold,
			TotalPriceSold:       product.TotalPriceSold,
		})
	}

	// save to database
	if err := s.invoiceRepo.Create(ctx, invoice); err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}

	return &models.CreateInvoiceResponse{
		InvoiceNo: invoice.InvoiceNo,
	}, nil
}
