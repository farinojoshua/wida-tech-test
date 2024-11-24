package invoices

import (
	"context"
	"fmt"
	"invoice-management-system/internal/models"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func (s *service) ImportInvoicesFromXLSX(ctx context.Context, file *multipart.FileHeader) (*models.ImportInvoiceResponse, error) {
	response := &models.ImportInvoiceResponse{
		ImportedCount: 0,
		Errors:        []models.ImportError{},
	}

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	// Parse XLSX
	xlsx, err := excelize.OpenReader(src)
	if err != nil {
		return nil, fmt.Errorf("failed to parse xlsx file: %w", err)
	}
	defer xlsx.Close()

	// Get both sheets
	invoiceSheet, err := xlsx.GetRows("invoice")
	if err != nil {
		return nil, fmt.Errorf("failed to read invoice sheet: %w", err)
	}

	productSheet, err := xlsx.GetRows("product sold")
	if err != nil {
		return nil, fmt.Errorf("failed to read product sold sheet: %w", err)
	}

	// First pass: validate and collect all data
	processedInvoices := make(map[string]bool)
	validInvoices := make([]*models.Invoice, 0)
	productsByInvoice := make(map[string][]models.Product)

	// Process products first
	for i, row := range productSheet {
		if i == 0 { // Skip header
			continue
		}
		if len(row) < 5 {
			response.Errors = append(response.Errors, models.ImportError{
				InvoiceNo: fmt.Sprintf("Row %d", i+1),
				Message:   "incomplete product data",
			})
			continue
		}

		invoiceNo := strings.TrimSpace(row[0])
		product := models.Product{
			InvoiceNo:            invoiceNo,
			ItemName:             strings.TrimSpace(row[1]),
			Quantity:             parseIntOrZero(row[2]),
			TotalCostOfGoodsSold: parseFloatOrZero(row[3]),
			TotalPriceSold:       parseFloatOrZero(row[4]),
		}

		// Validate product
		if err := validateProduct(product); err != nil {
			response.Errors = append(response.Errors, models.ImportError{
				InvoiceNo: invoiceNo,
				Message:   err.Error(),
			})
			continue
		}

		productsByInvoice[invoiceNo] = append(productsByInvoice[invoiceNo], product)
	}

	// Process invoices
	for i, row := range invoiceSheet {
		if i == 0 { // Skip header
			continue
		}
		if len(row) < 5 {
			continue
		}

		invoiceNo := strings.TrimSpace(row[0])

		// Check for duplicates within file
		if processedInvoices[invoiceNo] {
			response.Errors = append(response.Errors, models.ImportError{
				InvoiceNo: invoiceNo,
				Message:   "duplicate invoice number within file",
			})
			continue
		}

		// Check if invoice exists in database
		existing, err := s.invoiceRepo.GetByInvoiceNo(ctx, invoiceNo)
		if err != nil {
			return nil, fmt.Errorf("failed to check existing invoice: %w", err)
		}
		if existing != nil {
			response.Errors = append(response.Errors, models.ImportError{
				InvoiceNo: invoiceNo,
				Message:   "invoice number already exists in database",
			})
			continue
		}

		// Parse and validate invoice
		invoice, err := createInvoiceFromRow(row, productsByInvoice[invoiceNo])
		if err != nil {
			response.Errors = append(response.Errors, models.ImportError{
				InvoiceNo: invoiceNo,
				Message:   err.Error(),
			})
			continue
		}

		processedInvoices[invoiceNo] = true
		validInvoices = append(validInvoices, invoice)
	}

	// If there are any errors, don't save anything
	if len(response.Errors) > 0 {
		return response, nil
	}

	// Save valid invoices
	for _, invoice := range validInvoices {
		if err := s.invoiceRepo.Create(ctx, invoice); err != nil {
			return nil, fmt.Errorf("failed to save invoice %s: %w", invoice.InvoiceNo, err)
		}
		response.ImportedCount++
	}

	return response, nil
}

// Helper functions
func validateProduct(product models.Product) error {
	if len(product.ItemName) < 5 {
		return fmt.Errorf("item name must be at least 5 characters")
	}
	if product.Quantity < 1 {
		return fmt.Errorf("quantity must be at least 1")
	}
	if product.TotalCostOfGoodsSold < 0 {
		return fmt.Errorf("total cost of goods sold must be non-negative")
	}
	if product.TotalPriceSold < 0 {
		return fmt.Errorf("total price sold must be non-negative")
	}
	return nil
}

func createInvoiceFromRow(row []string, products []models.Product) (*models.Invoice, error) {
	// Get date string and clean it
	dateStr := strings.TrimSpace(row[1])

	// Split by hyphen
	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid date format: must be DD-MM-YY")
	}

	// Convert 2-digit year to 4-digit year
	day := parts[0]
	month := parts[1]
	year := parts[2]

	// Add leading zeros if needed
	if len(day) == 1 {
		day = "0" + day
	}
	if len(month) == 1 {
		month = "0" + month
	}

	// Convert 2-digit year to 4-digit year
	if len(year) == 2 {
		// Assume years 00-49 are 2000s, 50-99 are 1900s
		yearNum, err := strconv.Atoi(year)
		if err != nil {
			return nil, fmt.Errorf("invalid year format")
		}
		if yearNum < 50 {
			year = "20" + year
		} else {
			year = "19" + year
		}
	}

	// Reconstruct date in standard format
	standardDate := fmt.Sprintf("%s/%s/%s", day, month, year)

	// Parse the standardized date
	date, err := time.Parse("02/01/2006", standardDate)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}

	// Validate names
	customerName := strings.TrimSpace(row[2])
	if len(customerName) < 2 {
		return nil, fmt.Errorf("customer name must be at least 2 characters")
	}

	salespersonName := strings.TrimSpace(row[3])
	if len(salespersonName) < 2 {
		return nil, fmt.Errorf("salesperson name must be at least 2 characters")
	}

	// Validate payment type
	paymentType := strings.ToUpper(strings.TrimSpace(row[4]))
	if paymentType != string(models.PaymentTypeCash) && paymentType != string(models.PaymentTypeCredit) {
		return nil, fmt.Errorf("payment type must be either CASH or CREDIT")
	}

	// Create invoice
	invoice := &models.Invoice{
		InvoiceNo:       strings.TrimSpace(row[0]),
		Date:            date,
		CustomerName:    customerName,
		SalespersonName: salespersonName,
		PaymentType:     models.PaymentType(paymentType),
		Products:        products,
	}

	// Set notes if provided
	if len(row) > 5 && row[5] != "" {
		notes := strings.TrimSpace(row[5])
		if len(notes) < 5 {
			return nil, fmt.Errorf("notes must be at least 5 characters")
		}
		invoice.Notes = &notes
	}

	return invoice, nil
}

func parseIntOrZero(s string) int {
	var result int
	_, err := fmt.Sscanf(strings.TrimSpace(s), "%d", &result)
	if err != nil {
		return 0
	}
	return result
}

func parseFloatOrZero(s string) float64 {
	var result float64
	_, err := fmt.Sscanf(strings.TrimSpace(s), "%f", &result)
	if err != nil {
		return 0
	}
	return result
}
