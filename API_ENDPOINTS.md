# E-Commerce API Endpoints Documentation

## Authentication Endpoints

### Register
- **URL**: `http://localhost:8080/auth/register`
- **Method**: POST
- **Headers**: 
  - `Content-Type: application/json`
- **Body**:
```json
{
    "email": "ornek@email.com",
    "password": "123456",
    "first_name": "Ad",
    "last_name": "Soyad"
}
```
- **Success Response**: 201 Created
```json
{
    "id": 1,
    "email": "ornek@email.com",
    "first_name": "Ad",
    "last_name": "Soyad",
    "role": "user",
    "created_at": "2024-05-03T14:45:00Z",
    "updated_at": "2024-05-03T14:45:00Z"
}
```

### Login
- **URL**: `http://localhost:8080/auth/login`
- **Method**: POST
- **Headers**: 
  - `Content-Type: application/json`
- **Body**:
```json
{
    "email": "ornek@email.com",
    "password": "123456"
}
```
- **Success Response**: 200 OK
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
        "id": 1,
        "email": "ornek@email.com",
        "first_name": "Ad",
        "last_name": "Soyad",
        "role": "user"
    }
}
```

## Product Endpoints

### List Products
- **URL**: `http://localhost:8080/products`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}` (optional)
- **Success Response**: 200 OK
```json
[
    {
        "id": 1,
        "name": "Ürün Adı",
        "description": "Ürün Açıklaması",
        "price": 99.99,
        "stock": 100,
        "category": "Kategori",
        "image_url": "https://example.com/image.jpg",
        "sku": "PRD001",
        "is_active": true
    }
]
```

### Search Products
- **URL**: `http://localhost:8080/products/search?q={search_term}`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}` (optional)
- **Success Response**: 200 OK
```json
[
    {
        "id": 1,
        "name": "Ürün Adı",
        "description": "Ürün Açıklaması",
        "price": 99.99,
        "stock": 100,
        "category": "Kategori",
        "image_url": "https://example.com/image.jpg",
        "sku": "PRD001",
        "is_active": true
    }
]
```

### Get Product
- **URL**: `http://localhost:8080/products/{id}`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}` (optional)
- **Success Response**: 200 OK
```json
{
    "id": 1,
    "name": "Ürün Adı",
    "description": "Ürün Açıklaması",
    "price": 99.99,
    "stock": 100,
    "category": "Kategori",
    "image_url": "https://example.com/image.jpg",
    "sku": "PRD001",
    "is_active": true
}
```

### Create Product (Protected)
- **URL**: `http://localhost:8080/products`
- **Method**: POST
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "name": "Yeni Ürün",
    "description": "Ürün Açıklaması",
    "price": 99.99,
    "stock": 100,
    "category": "Kategori",
    "image_url": "https://example.com/image.jpg",
    "sku": "PRD001"
}
```

### Update Product (Protected)
- **URL**: `http://localhost:8080/products/{id}`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "name": "Güncellenmiş Ürün",
    "description": "Güncellenmiş Açıklama",
    "price": 89.99,
    "stock": 50,
    "category": "Yeni Kategori"
}
```

### Delete Product (Protected)
- **URL**: `http://localhost:8080/products/{id}`
- **Method**: DELETE
- **Headers**: 
  - `Authorization: Bearer {token}`

### Update Stock (Protected)
- **URL**: `http://localhost:8080/products/{id}/stock`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "stock": 150
}
```

## Order Endpoints (All Protected)

### Create Order
- **URL**: `http://localhost:8080/orders`
- **Method**: POST
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "items": [
        {
            "product_id": 1,
            "quantity": 2
        }
    ],
    "shipping_address": "Teslimat Adresi",
    "billing_address": "Fatura Adresi",
    "payment_method": "credit_card"
}
```

### Get Order
- **URL**: `http://localhost:8080/orders/{id}`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}`

### Get User Orders
- **URL**: `http://localhost:8080/orders/user/{userID}`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}`

### Update Order Status
- **URL**: `http://localhost:8080/orders/{id}/status`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "status": "processing"
}
```

### Cancel Order
- **URL**: `http://localhost:8080/orders/{id}`
- **Method**: DELETE
- **Headers**: 
  - `Authorization: Bearer {token}`

## Payment Endpoints (All Protected)

### Create Payment
- **URL**: `http://localhost:8080/payments`
- **Method**: POST
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "order_id": 1,
    "amount": 199.98,
    "payment_method": "credit_card"
}
```

### Process Payment
- **URL**: `http://localhost:8080/payments/{id}/process`
- **Method**: POST
- **Headers**: 
  - `Authorization: Bearer {token}`

### Get Payment
- **URL**: `http://localhost:8080/payments/{id}`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}`

### Get User Payments
- **URL**: `http://localhost:8080/payments/user/{user_id}`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}`

### Get Order Payments
- **URL**: `http://localhost:8080/payments/order/{order_id}`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}`

### Refund Payment
- **URL**: `http://localhost:8080/payments/{id}/refund`
- **Method**: POST
- **Headers**: 
  - `Authorization: Bearer {token}`

### List Payments
- **URL**: `http://localhost:8080/payments`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}`

## Notes
1. Tüm protected endpoint'ler için `Authorization` header'ında geçerli bir JWT token gereklidir.
2. Token formatı: `Bearer {token}`
3. Başarılı login sonrası alınan token'ı diğer isteklerde kullanabilirsiniz.
4. Tüm endpoint'lerde hata durumunda uygun HTTP status code'ları ve hata mesajları döner. 