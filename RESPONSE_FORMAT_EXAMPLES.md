# 🔹 Updated Response Format Examples

## ✅ Success Responses

### GET All Products
```json
{
  "success": true,
  "data": [
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
  ],
  "meta": {
    "total": 2,
    "page": 1,
    "limit": 2
  }
}
```

### GET Single Product
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "iPhone 15 Pro",
    "description": "Latest Apple iPhone",
    "shortDescription": "Flagship phone",
    "status": "active",
    "createdAt": "2025-08-17T05:39:06.351Z",
    "updatedAt": "2025-08-17T05:39:06.351Z"
  }
}
```

### POST Create Product
```json
{
  "success": true,
  "data": {
    "id": 3,
    "name": "iPad Pro",
    "description": "Professional tablet with M2 chip",
    "shortDescription": "iPad Pro with M2",
    "status": "active",
    "createdAt": "2025-08-17T05:45:30.123Z",
    "updatedAt": "2025-08-17T05:45:30.123Z"
  }
}
```

## ❌ Error Responses

### Validation Error (400)
```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Key: 'Product.Name' Error:Field validation for 'Name' failed on the 'required' tag"
  }
}
```

### Not Found Error (404)
```json
{
  "success": false,
  "error": {
    "code": "NOT_FOUND",
    "message": "Product not found"
  }
}
```

### Bad Request Error (400)
```json
{
  "success": false,
  "error": {
    "code": "BAD_REQUEST",
    "message": "Invalid product ID format"
  }
}
```

### Internal Server Error (500)
```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "Database connection failed"
  }
}
```

## 🔧 Testing Examples

### Test with cURL

#### Create Product
```bash
curl -X POST http://localhost:8080/api/v1/products/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "MacBook Pro",
    "description": "Professional laptop with M3 chip",
    "shortDescription": "MacBook Pro M3",
    "status": "active"
  }' | jq
```

#### Get All Products
```bash
curl -X GET http://localhost:8080/api/v1/products/ | jq
```

#### Test Error Response
```bash
curl -X GET http://localhost:8080/api/v1/products/invalid_id | jq
```

## 📊 Response Structure

All API responses now follow this standardized format:

```go
type Response struct {
    Success bool        `json:"success"`           // Always present
    Data    interface{} `json:"data,omitempty"`    // Present on success
    Error   *ErrorInfo  `json:"error,omitempty"`   // Present on error
    Meta    interface{} `json:"meta,omitempty"`    // Optional metadata
}

type ErrorInfo struct {
    Code    string `json:"code"`     // Error code (e.g., "VALIDATION_ERROR")
    Message string `json:"message"`  // Human-readable error message
}
```

## 🎯 Benefits

- ✅ **Consistent Structure** - All responses follow the same format
- ✅ **Clear Success/Error Indication** - `success` field makes status obvious
- ✅ **Structured Error Handling** - Error codes and messages are standardized
- ✅ **Metadata Support** - Pagination and other metadata can be included
- ✅ **Client-Friendly** - Easy to parse and handle in frontend applications
