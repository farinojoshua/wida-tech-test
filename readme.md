# Technical Challenge Solution

This repository contains the solution for the Back End Engineer Technical Challenge, consisting of three main sections:

1. Invoice Management System (Section 1 & 2) - A RESTful API service built with Go
2. Algorithm Solution (Section 3) - Implementation of non-repeating digit combination problem

## Invoice Management System (Section 1 & 2)

A RESTful API service for managing invoices built with Go. This system allows for basic CRUD operations on invoices and their associated products, as well as bulk importing of invoice data via XLSX files.

## Features

- CRUD operations for invoices
- Import invoices from XLSX files
- Database migrations
- Docker support for MySQL database
- Input validation
- Error handling
- Transaction management

## Project Structure

According to the technical challenge requirements, this repository is organized as follows:

```
.
├── Algorithm/                         # Section 3 solution
│   └── solution.go                   # Non-repeating digit combination problem
├── Go/                                  # Section 1 & 2 source code
│   └── invoice-management-system/       # Invoice Management System Implementation
│       ├── cmd/
│       │   └── main.go                 # Application entry point
│       ├── internal/
│       │   ├── configs/                # Configuration management
│       │   ├── handlers/               # HTTP request handlers
│       │   ├── models/                 # Data models and validation
│       │   ├── repository/             # Database operations
│       │   └── service/                # Business logic
│       ├── pkg/
│       │   └── internalsql/           # Database connection utilities
│       ├── scripts/
│       │   └── migrations/            # Database migration files
│       ├── docker-compose.yml         # Docker configuration
│       └── Makefile                  # Build and deployment commands
└── README.md                        # Project documentation & API collection
```

## Prerequisites

- Go 1.19 or higher
- Docker and Docker Compose
- Make

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/farinojoshua/wida-tech-test.git
   cd Go/invoice-management-system
   ```

2. Start the MySQL database using Docker:

   ```bash
   docker-compose up -d
   ```

3. Run database migrations:

   ```bash
   make migrate-up
   ```

4. Run the application:
   ```bash
   go run cmd/main.go
   ```

## Postman Collection

Below is the Postman collection v2.1 JSON for testing the APIs:

````json
{
  "info": {
    "_postman_id": "bd70e18a-c62b-4fbf-8ccc-90077ae52221",
    "name": "Invoice Management System",
    "description": "# 🚀 Get started here\n\nThis template guides you through CRUD operations (GET, POST, PUT, DELETE), variables, and tests.\n\n## 🔖 **How to use this template**\n\n#### **Step 1: Send requests**\n\nRESTful APIs allow you to perform CRUD operations using the POST, GET, PUT, and DELETE HTTP methods.\n\nThis collection contains each of these [request](https://learning.postman.com/docs/sending-requests/requests/) types. Open each request and click \"Send\" to see what happens.\n\n#### **Step 2: View responses**\n\nObserve the response tab for status code (200 OK), response time, and size.\n\n#### **Step 3: Send new Body data**\n\nUpdate or add new data in \"Body\" in the POST request. Typically, Body data is also used in PUT request.\n\n```\n{\n    \"name\": \"Add your name in the body\"\n}\n\n ```\n\n#### **Step 4: Update the variable**\n\nVariables enable you to store and reuse values in Postman. We have created a [variable](https://learning.postman.com/docs/sending-requests/variables/) called `base_url` with the sample request [https://postman-api-learner.glitch.me](https://postman-api-learner.glitch.me). Replace it with your API endpoint to customize this collection.\n\n#### **Step 5: Add tests in the \"Scripts\" tab**\n\nAdding tests to your requests can help you confirm that your API is working as expected. You can write test scripts in JavaScript and view the output in the \"Test Results\" tab.\n\n<img src=\"https://content.pstmn.io/fa30ea0a-373d-4545-a668-e7b283cca343/aW1hZ2UucG5n\" width=\"2162\" height=\"1530\">\n\n## 💪 Pro tips\n\n- Use folders to group related requests and organize the collection.\n    \n- Add more [scripts](https://learning.postman.com/docs/writing-scripts/intro-to-scripts/) to verify if the API works as expected and execute workflows.\n    \n\n## 💡Related templates\n\n[API testing basics](https://go.postman.co/redirect/workspace?type=personal&collectionTemplateId=e9a37a28-055b-49cd-8c7e-97494a21eb54&sourceTemplateId=ddb19591-3097-41cf-82af-c84273e56719)  \n[API documentation](https://go.postman.co/redirect/workspace?type=personal&collectionTemplateId=e9c28f47-1253-44af-a2f3-20dce4da1f18&sourceTemplateId=ddb19591-3097-41cf-82af-c84273e56719)  \n[Authorization methods](https://go.postman.co/redirect/workspace?type=personal&collectionTemplateId=31a9a6ed-4cdf-4ced-984c-d12c9aec1c27&sourceTemplateId=ddb19591-3097-41cf-82af-c84273e56719)",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
    "_exporter_id": "18887807"
  },
  "item": [
    {
      "name": "Get data",
      "event": [
        {
          "listen": "test",
          "script": {
            "exec": [
              "pm.test(\"Status code is 200\", function () {",
              "    pm.response.to.have.status(200);",
              "});"
            ],
            "type": "text/javascript",
            "packages": {}
          }
        }
      ],
      "request": {
        "method": "GET",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": "",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "{{base_url}}/api/v1/invoices?start_date=2021-01-01&end_date=2024-11-22&page=3&size=3",
          "host": ["{{base_url}}"],
          "path": ["api", "v1", "invoices"],
          "query": [
            {
              "key": "start_date",
              "value": "2021-01-01"
            },
            {
              "key": "end_date",
              "value": "2024-11-22"
            },
            {
              "key": "page",
              "value": "3"
            },
            {
              "key": "size",
              "value": "3"
            }
          ]
        },
        "description": "This is a GET request and it is used to \"get\" data from an endpoint. There is no request body for a GET request, but you can use query parameters to help specify the resource you want data on (e.g., in this request, we have `id=1`).\n\nA successful GET response will have a `200 OK` status, and should include some kind of response body - for example, HTML web content or JSON data."
      },
      "response": []
    },
    {
      "name": "Post data",
      "event": [
        {
          "listen": "test",
          "script": {
            "exec": [
              "pm.test(\"Successful POST request\", function () {",
              "    pm.expect(pm.response.code).to.be.oneOf([200, 201]);",
              "});",
              ""
            ],
            "type": "text/javascript",
            "packages": {}
          }
        }
      ],
      "request": {
        "method": "POST",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"invoice_no\": \"INV-DSA\",\n  \"date\": \"2024-04-22\",\n  \"customer_name\": \"Farino\",\n  \"salesperson_name\": \"Jane Doe\",\n  \"payment_type\": \"CASH\",\n  \"notes\": \"First invoice\",\n  \"products\": [\n    {\n      \"item_name\": \"Product A\",\n      \"quantity\": 2,\n      \"total_cost_of_goods_sold\": 100.00,\n      \"total_price_sold\": 150.00\n    },\n    {\n      \"item_name\": \"Product B\",\n      \"quantity\": 1,\n      \"total_cost_of_goods_sold\": 200.00,\n      \"total_price_sold\": 250.00\n    }\n  ]\n}\n",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "{{base_url}}/api/v1/invoices",
          "host": ["{{base_url}}"],
          "path": ["api", "v1", "invoices"]
        },
        "description": "This is a POST request, submitting data to an API via the request body. This request submits JSON data, and the data is reflected in the response.\n\nA successful POST request typically returns a `200 OK` or `201 Created` response code."
      },
      "response": []
    },
    {
      "name": "Update data",
      "event": [
        {
          "listen": "test",
          "script": {
            "exec": [
              "pm.test(\"Successful PUT request\", function () {",
              "    pm.expect(pm.response.code).to.be.oneOf([200, 201, 204]);",
              "});",
              ""
            ],
            "type": "text/javascript",
            "packages": {}
          }
        }
      ],
      "request": {
        "method": "PUT",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"invoice_no\": \"INV-DSA\",\n  \"date\": \"2023-11-21\",\n  \"customer_name\": \"Updated Customer\",\n  \"salesperson_name\": \"Updated Salesperson\",\n  \"payment_type\": \"CREDIT\",\n  \"notes\": \"Updated notes\",\n  \"products\": [\n    {\n      \"item_name\": \"Updated Product A\",\n      \"quantity\": 2,\n      \"total_cost_of_goods_sold\": 200.00,\n      \"total_price_sold\": 300.00\n    },\n    {\n      \"item_name\": \"Updated Product B\",\n      \"quantity\": 3,\n      \"total_cost_of_goods_sold\": 300.00,\n      \"total_price_sold\": 450.00\n    }\n  ]\n}",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "{{base_url}}/api/v1/invoices/INV-DSA",
          "host": ["{{base_url}}"],
          "path": ["api", "v1", "invoices", "INV-DSA"]
        },
        "description": "This is a PUT request and it is used to overwrite an existing piece of data. For instance, after you create an entity with a POST request, you may want to modify that later. You can do that using a PUT request. You typically identify the entity being updated by including an identifier in the URL (eg. `id=1`).\n\nA successful PUT request typically returns a `200 OK`, `201 Created`, or `204 No Content` response code."
      },
      "response": []
    },
    {
      "name": "Delete data",
      "event": [
        {
          "listen": "test",
          "script": {
            "exec": [
              "pm.test(\"Successful DELETE request\", function () {",
              "    pm.expect(pm.response.code).to.be.oneOf([200, 202, 204]);",
              "});",
              ""
            ],
            "type": "text/javascript",
            "packages": {}
          }
        }
      ],
      "request": {
        "method": "DELETE",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": "",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "{{base_url}}/api/v1/invoices/INV-DSA",
          "host": ["{{base_url}}"],
          "path": ["api", "v1", "invoices", "INV-DSA"]
        },
        "description": "This is a DELETE request, and it is used to delete data that was previously created via a POST request. You typically identify the entity being updated by including an identifier in the URL (eg. `id=1`).\n\nA successful DELETE request typically returns a `200 OK`, `202 Accepted`, or `204 No Content` response code."
      },
      "response": []
    },
    {
      "name": "Import Excel",
      "request": {
        "method": "POST",
        "header": [],
        "body": {
          "mode": "formdata",
          "formdata": [
            {
              "key": "file",
              "type": "file",
              "src": "/D:/kerja/Daftar Kerja/Wida Tech/Technical Test/InvoiceImport.xlsx"
            }
          ]
        },
        "url": {
          "raw": "{{base_url}}/api/v1/invoices/import",
          "host": ["{{base_url}}"],
          "path": ["api", "v1", "invoices", "import"]
        }
      },
      "response": []
    }
  ],
  "event": [
    {
      "listen": "prerequest",
      "script": {
        "type": "text/javascript",
        "exec": [""]
      }
    },
    {
      "listen": "test",
      "script": {
        "type": "text/javascript",
        "exec": [""]
      }
    }
  ],
  "variable": [
    {
      "key": "id",
      "value": "1"
    },
    {
      "key": "base_url",
      "value": "http://localhost:8080"
    }
  ]
}
````

## API Documentation

### 1. Create Invoice

- **URL**: `/api/v1/invoices`
- **Method**: `POST`
- **Request Body**:

```json
{
  "invoice_no": "INV001",
  "date": "2024-01-01",
  "customer_name": "John Doe",
  "salesperson_name": "Jane Smith",
  "payment_type": "CASH",
  "notes": "Optional notes",
  "products": [
    {
      "item_name": "Product A",
      "quantity": 2,
      "total_cost_of_goods_sold": 100.0,
      "total_price_sold": 150.0
    }
  ]
}
```

### 2. Get Invoices

- **URL**: `/api/v1/invoices`
- **Method**: `GET`
- **Query Parameters**:
  - `start_date`: YYYY-MM-DD (required)
  - `end_date`: YYYY-MM-DD (required)
  - `page`: int (required, min: 1)
  - `size`: int (required, min: 1)

### 3. Update Invoice

- **URL**: `/api/v1/invoices/:invoice_no`
- **Method**: `PUT`
- **Request Body**: Same as Create Invoice

### 4. Delete Invoice

- **URL**: `/api/v1/invoices/:invoice_no`
- **Method**: `DELETE`

### 5. Import Invoices

- **URL**: `/api/v1/invoices/import`
- **Method**: `POST`
- **Content-Type**: `multipart/form-data`
- **Form Parameters**:
  - `file`: XLSX file with two sheets (invoice and product sold)

## XLSX Import Format

### Invoice Sheet

Required columns:

- Invoice No
- Date (DD-MM-YY)
- Customer Name
- Salesperson Name
- Payment Type
- Notes (optional)

### Product Sold Sheet

Required columns:

- Invoice No
- Item Name
- Quantity
- Total Cost of Goods Sold
- Total Price Sold

## Error Handling

The API returns standardized error responses in the following format:

```json
{
  "status": "error",
  "message": "Error description",
  "errors": [
    {
      "field": "field_name",
      "message": "Error message"
    }
  ]
}
```

## Data Validation Rules

### Invoice

- `invoice_no`: Required, unique
- `date`: Required, valid date format
- `customer_name`: Required, minimum 2 characters
- `salesperson_name`: Required, minimum 2 characters
- `payment_type`: Required, either "CASH" or "CREDIT"
- `notes`: Optional, minimum 5 characters if provided
- `products`: At least one product required

### Product

- `item_name`: Required, minimum 5 characters
- `quantity`: Required, minimum 1
- `total_cost_of_goods_sold`: Required, non-negative
- `total_price_sold`: Required, non-negative

## Development

### Creating New Migrations

```bash
make migrate-create name=migration_name
```

### Rolling Back Migrations

```bash
make migrate-down
```

## Algorithm Solution (Section 3)

Implementation of a function that finds all possible combinations of non-repeating digits (1-9) given length `l` and target sum `t`.

### Problem Description

Write the most efficient function that returns a list of all possible combinations of non-repeating digit (1-9) given variable `l` and `t`.

- `l` is the length of a combination
- `t` is the total of all numbers in the combination

### Rules and Constraints

1. **Digit Range**:

   - Only use numbers from 1 to 9
   - Numbers cannot be repeated in a combination

2. **Combination Rules**:
   - Each number can only be used once
   - Order doesn't matter ([1,2,3] is considered the same as [3,2,1])
   - Only one version of each combination should appear in the result

### Solution Implementation

#### Approach

The solution uses backtracking with several optimizations:

1. **Backtracking Strategy**:

   - Systematically builds combinations using recursion
   - Prunes invalid branches early
   - Maintains sorted order to avoid duplicates

2. **Main Components**:
   - `findCombinations`: Main entry point function
   - `backtrack`: Recursive helper function for combination building

The implementation passes all test cases and follows the problem constraints.
