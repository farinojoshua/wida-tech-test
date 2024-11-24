package invoices

import (
	"invoice-management-system/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInvoices(c *gin.Context) {
	var req models.GetInvoicesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Invalid query parameters",
			Errors: []models.Error{{
				Field:   "query",
				Message: err.Error(),
			}},
		})
		return
	}

	// Set default values if not provided
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	resp, err := h.invoiceService.GetInvoices(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "error",
			Message: "Failed to fetch invoices",
			Errors: []models.Error{{
				Field:   "server",
				Message: err.Error(),
			}},
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Data:   resp,
	})
}
