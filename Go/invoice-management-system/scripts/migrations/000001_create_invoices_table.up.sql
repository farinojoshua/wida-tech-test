CREATE TABLE IF NOT EXISTS invoices (
    invoice_no VARCHAR(255) PRIMARY KEY,
    date DATE NOT NULL,
    customer_name VARCHAR(255) NOT NULL CHECK (LENGTH(customer_name) >= 2),
    salesperson_name VARCHAR(255) NOT NULL CHECK (LENGTH(salesperson_name) >= 2),
    payment_type ENUM('CASH', 'CREDIT') NOT NULL,
    notes TEXT CHECK (notes IS NULL OR LENGTH(notes) >= 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);