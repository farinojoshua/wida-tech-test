package models

import (
	"fmt"
	"time"
)

func (i *CreateInvoiceRequest) Validate() []Error {
	var errors []Error

	// Validate date format
	_, err := time.Parse("2006-01-02", i.Date)
	if err != nil {
		errors = append(errors, Error{
			Field:   "date",
			Message: "invalid date format, must be YYYY-MM-DD",
		})
	}

	// Validate notes if provided
	if i.Notes != nil && len(*i.Notes) < 5 {
		errors = append(errors, Error{
			Field:   "notes",
			Message: "notes must be at least 5 characters long",
		})
	}

	// Validate products
	if len(i.Products) == 0 {
		errors = append(errors, Error{
			Field:   "products",
			Message: "at least one product is required",
		})
	}

	for idx, product := range i.Products {
		if len(product.ItemName) < 5 {
			errors = append(errors, Error{
				Field:   fmt.Sprintf("products[%d].item_name", idx),
				Message: "item name must be at least 5 characters long",
			})
		}
		if product.Quantity < 1 {
			errors = append(errors, Error{
				Field:   fmt.Sprintf("products[%d].quantity", idx),
				Message: "quantity must be at least 1",
			})
		}
		if product.TotalCostOfGoodsSold < 0 {
			errors = append(errors, Error{
				Field:   fmt.Sprintf("products[%d].total_cost_of_goods_sold", idx),
				Message: "total cost of goods sold must be non-negative",
			})
		}
		if product.TotalPriceSold < 0 {
			errors = append(errors, Error{
				Field:   fmt.Sprintf("products[%d].total_price_sold", idx),
				Message: "total price sold must be non-negative",
			})
		}
	}

	return errors
}

func (i *UpdateInvoiceRequest) Validate() []Error {
	var errors []Error

	// Validate date format
	_, err := time.Parse("2006-01-02", i.Date)
	if err != nil {
		errors = append(errors, Error{
			Field:   "date",
			Message: "invalid date format, must be YYYY-MM-DD",
		})
	}

	// Validate notes if provided
	if i.Notes != nil && len(*i.Notes) < 5 {
		errors = append(errors, Error{
			Field:   "notes",
			Message: "notes must be at least 5 characters long",
		})
	}

	// Validate products
	if len(i.Products) == 0 {
		errors = append(errors, Error{
			Field:   "products",
			Message: "at least one product is required",
		})
	}

	for idx, product := range i.Products {
		if len(product.ItemName) < 5 {
			errors = append(errors, Error{
				Field:   fmt.Sprintf("products[%d].item_name", idx),
				Message: "item name must be at least 5 characters long",
			})
		}
		if product.Quantity < 1 {
			errors = append(errors, Error{
				Field:   fmt.Sprintf("products[%d].quantity", idx),
				Message: "quantity must be at least 1",
			})
		}
		if product.TotalCostOfGoodsSold < 0 {
			errors = append(errors, Error{
				Field:   fmt.Sprintf("products[%d].total_cost_of_goods_sold", idx),
				Message: "total cost of goods sold must be non-negative",
			})
		}
		if product.TotalPriceSold < 0 {
			errors = append(errors, Error{
				Field:   fmt.Sprintf("products[%d].total_price_sold", idx),
				Message: "total price sold must be non-negative",
			})
		}
	}

	return errors
}
