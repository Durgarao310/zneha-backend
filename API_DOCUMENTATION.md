# 📖 Zneha Backend API Documentation

## Overview
This is the complete API documentation for the Zneha Backend Product Management System.

## Base URL
```
http://localhost:8080
```

## API Version
```
v1
```

## Endpoints

### Products API

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/products/` | Create a new product |
| GET | `/api/v1/products/` | Get all products |
| GET | `/api/v1/products/:id` | Get product by ID |
| PUT | `/api/v1/products/:id` | Update product |
| DELETE | `/api/v1/products/:id` | Delete product |

---

## 📝 Product Model

```json
{
    "id": 1,
    "name": "Product Name",
    "description": "Detailed product description",
    "shortDescription": "Brief description",
    "status": "active",
    "createdAt": "2025-08-17T05:39:06.351Z",
    "updatedAt": "2025-08-17T05:39:06.351Z"
}
```

### Field Descriptions

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | `uint64` | Auto | Unique product identifier |
| `name` | `string` | ✅ | Product name (max 255 chars) |
| `description` | `string` | ❌ | Detailed product description |
| `shortDescription` | `string` | ❌ | Brief product summary |
| `status` | `string` | ❌ | Product status (`active`, `inactive`) |
| `createdAt` | `timestamp` | Auto | Creation timestamp |
| `updatedAt` | `timestamp` | Auto | Last update timestamp |

---

## 🚀 API Endpoints Details

### 1. Create Product
**POST** `/api/v1/products/`

**Request Body:**
```json
{
    "name": "iPhone 15 Pro Max",
    "description": "The most advanced iPhone with A17 Pro chip",
    "shortDescription": "Latest iPhone Pro Max",
    "status": "active"
}
```

**Response (201):**
```json
{
    "id": 1,
    "name": "iPhone 15 Pro Max",
    "description": "The most advanced iPhone with A17 Pro chip",
    "shortDescription": "Latest iPhone Pro Max",
    "status": "active",
    "createdAt": "2025-08-17T05:45:30.123Z",
    "updatedAt": "2025-08-17T05:45:30.123Z"
}
```

### 2. Get All Products
**GET** `/api/v1/products/`

**Response (200):**
```json
[
    {
        "id": 1,
        "name": "iPhone 15",
        "description": "Latest Apple iPhone",
        "shortDescription": "Apple flagship",
        "status": "active",
        "createdAt": "2025-08-17T05:39:06.351Z",
        "updatedAt": "2025-08-17T05:39:06.351Z"
    },
    {
        "id": 2,
        "name": "Samsung Galaxy S24",
        "description": "Flagship Android phone",
        "shortDescription": "Samsung flagship",
        "status": "active",
        "createdAt": "2025-08-17T05:39:06.352Z",
        "updatedAt": "2025-08-17T05:39:06.352Z"
    }
]
```

### 3. Get Product by ID
**GET** `/api/v1/products/:id`

**Response (200):**
```json
{
    "id": 1,
    "name": "iPhone 15",
    "description": "Latest Apple iPhone",
    "shortDescription": "Apple flagship",
    "status": "active",
    "createdAt": "2025-08-17T05:39:06.351Z",
    "updatedAt": "2025-08-17T05:39:06.351Z"
}
```

**Error Response (404):**
```json
{
    "error": "Product not found"
}
```

### 4. Update Product
**PUT** `/api/v1/products/:id`

**Request Body:**
```json
{
    "name": "iPhone 15 Pro Max - Updated",
    "description": "Updated description with new features",
    "shortDescription": "Updated iPhone Pro Max",
    "status": "active"
}
```

**Response (200):**
```json
{
    "id": 1,
    "name": "iPhone 15 Pro Max - Updated",
    "description": "Updated description with new features",
    "shortDescription": "Updated iPhone Pro Max",
    "status": "active",
    "createdAt": "2025-08-17T05:39:06.351Z",
    "updatedAt": "2025-08-17T05:45:30.789Z"
}
```

### 5. Delete Product
**DELETE** `/api/v1/products/:id`

**Response (204):** No Content

---

## 📋 Status Codes

| Code | Description |
|------|-------------|
| 200 | OK - Request successful |
| 201 | Created - Resource created successfully |
| 204 | No Content - Resource deleted successfully |
| 400 | Bad Request - Invalid request data |
| 404 | Not Found - Resource not found |
| 500 | Internal Server Error - Server error |

---

## 🧪 Testing with Postman

### Import Collection
1. Download `postman-collection.json`
2. Open Postman
3. Click **Import** → **Upload Files**
4. Select the collection file
5. Import `postman-environment.json` for environment variables

### Environment Variables
- `base_url`: `http://localhost:8080`
- `api_version`: `v1`

### Running Tests
Each request includes automated tests that verify:
- Response status codes
- Response structure
- Data validation
- Performance (response time < 2000ms)

---

## 🔧 Sample cURL Commands

### Create Product
```bash
curl -X POST http://localhost:8080/api/v1/products/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPad Pro",
    "description": "Professional tablet with M2 chip",
    "shortDescription": "iPad Pro with M2",
    "status": "active"
  }'
```

### Get All Products
```bash
curl -X GET http://localhost:8080/api/v1/products/
```

### Get Product by ID
```bash
curl -X GET http://localhost:8080/api/v1/products/1
```

### Update Product
```bash
curl -X PUT http://localhost:8080/api/v1/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPad Pro - Updated",
    "description": "Updated professional tablet",
    "shortDescription": "Updated iPad Pro",
    "status": "active"
  }'
```

### Delete Product
```bash
curl -X DELETE http://localhost:8080/api/v1/products/1
```

---

## 🚀 Quick Start

1. **Start the server:**
   ```bash
   make run
   ```

2. **Import Postman Collection:**
   - Import `postman-collection.json`
   - Import `postman-environment.json`

3. **Test the API:**
   - Run the collection tests
   - Or use individual requests

4. **Verify Database:**
   ```bash
   make migrate  # Create tables
   make seed     # Add sample data
   ```

---

## 📊 Error Handling

All errors return JSON format:
```json
{
    "error": "Error description"
}
```

Common error scenarios:
- Invalid JSON format → 400 Bad Request
- Missing required fields → 400 Bad Request  
- Product not found → 404 Not Found
- Database connection issues → 500 Internal Server Error
