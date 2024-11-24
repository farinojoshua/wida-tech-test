package invoices

import (
	"context"
	"fmt"
	"invoice-management-system/internal/models"
	"time"
)

func (s *service) GetInvoices(ctx context.Context, req *models.GetInvoicesRequest) (*models.GetInvoicesResponse, error) {
	// Validate date format
	_, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	_, err = time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	// Validate date range
	startDate, _ := time.Parse("2006-01-02", req.StartDate)
	endDate, _ := time.Parse("2006-01-02", req.EndDate)
	if endDate.Before(startDate) {
		return nil, fmt.Errorf("end date cannot be before start date")
	}

	// Get invoices
	invoices, err := s.invoiceRepo.Get(ctx, req.StartDate, req.EndDate, req.Page, req.Size)
	if err != nil {
		return nil, fmt.Errorf("failed to get invoices: %w", err)
	}

	// Get metadata
	metadata, err := s.invoiceRepo.GetMetadata(ctx, req.StartDate, req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	// Calculate pagination info
	metadata.CurrentPage = req.Page
	metadata.TotalPages = (metadata.TotalRecords + req.Size - 1) / req.Size

	return &models.GetInvoicesResponse{
		Invoices: invoices,
		Metadata: *metadata,
	}, nil
}
