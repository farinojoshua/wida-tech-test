package invoices

import (
	"context"
	"invoice-management-system/internal/configs"
	"invoice-management-system/internal/models"
)

type invoiceRepositoy interface {
	Create(ctx context.Context, invoice *models.Invoice) error
	Get(ctx context.Context, startDate, endDate string, page, size int) ([]models.Invoice, error)
	GetByInvoiceNo(ctx context.Context, invoiceNo string) (*models.Invoice, error)
	Update(ctx context.Context, invoice *models.Invoice) error
	Delete(ctx context.Context, invoiceNo string) error
	GetMetadata(ctx context.Context, startDate, endDate string) (*models.Metadata, error)
}

type service struct {
	cfg         *configs.Config
	invoiceRepo invoiceRepositoy
}

func NewService(cfg *configs.Config, invoiceRepo invoiceRepositoy) *service {
	return &service{
		cfg:         cfg,
		invoiceRepo: invoiceRepo,
	}
}
