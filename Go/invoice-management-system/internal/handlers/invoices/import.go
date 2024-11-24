package invoices

import (
	"invoice-management-system/internal/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ImportInvoices(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "no file uploaded",
			Errors: []models.Error{{
				Field:   "file",
				Message: err.Error(),
			}},
		})
		return
	}

	// Validate file extension
	if !strings.HasSuffix(file.Filename, ".xlsx") {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "invalid file type",
			Errors: []models.Error{{
				Field:   "file",
				Message: "only XLSX files are allowed",
			}},
		})
		return
	}

	resp, err := h.invoiceService.ImportInvoicesFromXLSX(c.Request.Context(), file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "error",
			Message: "failed to process file",
			Errors: []models.Error{{
				Field:   "server",
				Message: err.Error(),
			}},
		})
		return
	}

	if len(resp.Errors) > 0 {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "validation errors found in file",
			Data:    resp,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "file imported successfully",
		Data:    resp,
	})
}
