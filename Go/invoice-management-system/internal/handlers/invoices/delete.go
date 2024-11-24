package invoices

import (
	"invoice-management-system/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteInvoice(c *gin.Context) {
	invoiceNo := c.Param("invoice_no")
	if invoiceNo == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Invoice number is required",
			Errors: []models.Error{{
				Field:   "invoice_no",
				Message: "invoice number is required",
			}},
		})
		return
	}

	err := h.invoiceService.DeleteInvoice(c.Request.Context(), invoiceNo)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "invoice not found" {
			statusCode = http.StatusNotFound
		}

		c.JSON(statusCode, models.Response{
			Status:  "error",
			Message: "Failed to delete invoice",
			Errors: []models.Error{{
				Field:   "server",
				Message: err.Error(),
			}},
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Invoice deleted successfully",
	})
}
