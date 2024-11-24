package invoices

import (
	"invoice-management-system/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateInvoice(c *gin.Context) {
	invoiceNo := c.Param("invoice_no")
	if invoiceNo == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Invalid invoice number",
			Errors: []models.Error{{
				Field:   "invoice_no",
				Message: "invoice number is required",
			}},
		})
		return
	}

	var req models.UpdateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Invalid request body",
			Errors: []models.Error{{
				Field:   "request body",
				Message: err.Error(),
			}},
		})
		return
	}

	err := h.invoiceService.UpdateInvoice(c.Request.Context(), invoiceNo, &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "invoice not found" {
			statusCode = http.StatusNotFound
		}

		c.JSON(statusCode, models.Response{
			Status:  "error",
			Message: "Failed to update invoice",
			Errors: []models.Error{{
				Field:   "server",
				Message: err.Error(),
			}},
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Invoice updated successfully",
	})
}
