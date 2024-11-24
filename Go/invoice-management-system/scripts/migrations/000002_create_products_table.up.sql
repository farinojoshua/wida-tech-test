CREATE TABLE IF NOT EXISTS products (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    invoice_no VARCHAR(255) NOT NULL,
    item_name VARCHAR(255) NOT NULL CHECK (LENGTH(item_name) >= 5),
    quantity INT NOT NULL CHECK (quantity >= 1),
    total_cost_of_goods_sold DECIMAL(10,2) NOT NULL CHECK (total_cost_of_goods_sold >= 0),
    total_price_sold DECIMAL(10,2) NOT NULL CHECK (total_price_sold >= 0),
    FOREIGN KEY (invoice_no) REFERENCES invoices(invoice_no) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
