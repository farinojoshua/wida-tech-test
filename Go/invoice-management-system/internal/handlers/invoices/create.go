package invoices

import (
	"invoice-management-system/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateInvoice(c *gin.Context) {
	var req models.CreateInvoiceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "invalid request body",
			Errors: []models.Error{{
				Field:   "request body",
				Message: err.Error(),
			}},
		})
		return
	}

	resp, err := h.invoiceService.CreateInvoice(c.Request.Context(), &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "invoice number already exists" {
			statusCode = http.StatusConflict
		}

		c.JSON(statusCode, models.Response{
			Status:  "error",
			Message: "failed to create invoice",
			Errors: []models.Error{{
				Field:   "invoice number",
				Message: err.Error(),
			}},
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Status:  "success",
		Message: "invoice created successfully",
		Data:    resp,
	})
}
