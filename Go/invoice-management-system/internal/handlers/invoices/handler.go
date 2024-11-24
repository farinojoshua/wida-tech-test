package invoices

import (
	"context"
	"invoice-management-system/internal/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type invoiceService interface {
	CreateInvoice(ctx context.Context, req *models.CreateInvoiceRequest) (*models.CreateInvoiceResponse, error)
	GetInvoices(ctx context.Context, req *models.GetInvoicesRequest) (*models.GetInvoicesResponse, error)
	UpdateInvoice(ctx context.Context, invoiceNo string, req *models.UpdateInvoiceRequest) error
	DeleteInvoice(ctx context.Context, invoiceNo string) error
	ImportInvoicesFromXLSX(ctx context.Context, file *multipart.FileHeader) (*models.ImportInvoiceResponse, error)
}

type Handler struct {
	*gin.Engine

	invoiceService invoiceService
}

func NewHandler(api *gin.Engine, invoiceService invoiceService) *Handler {
	return &Handler{
		Engine:         api,
		invoiceService: invoiceService,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/api/v1/invoices")
	{
		route.POST("", h.CreateInvoice)
		route.GET("", h.GetInvoices)
		route.PUT("/:invoice_no", h.UpdateInvoice)
		route.DELETE("/:invoice_no", h.DeleteInvoice)
		route.POST("/import", h.ImportInvoices)
	}
}
