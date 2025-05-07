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

## User Management Endpoints (All Protected)

### Get User Profile
- **URL**: `http://localhost:8080/users/{id}`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}`
- **Success Response**: 200 OK
```json
{
    "id": 1,
    "email": "ornek@email.com",
    "first_name": "Ad",
    "last_name": "Soyad",
    "role": "user",
    "is_active": true,
    "created_at": "2024-05-03T14:45:00Z",
    "updated_at": "2024-05-03T14:45:00Z"
}
```

### Update User Profile
- **URL**: `http://localhost:8080/users/{id}`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "first_name": "Yeni Ad",
    "last_name": "Yeni Soyad"
}
```
- **Success Response**: 200 OK
```json
{
    "id": 1,
    "email": "ornek@email.com",
    "first_name": "Yeni Ad",
    "last_name": "Yeni Soyad",
    "role": "user",
    "is_active": true,
    "created_at": "2024-05-03T14:45:00Z",
    "updated_at": "2024-05-03T14:45:00Z"
}
```

### Change Password
- **URL**: `http://localhost:8080/users/{id}/password`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "old_password": "eski123",
    "new_password": "yeni456"
}
```
- **Success Response**: 200 OK

### List Users (Admin Only)
- **URL**: `http://localhost:8080/users?page=1&limit=10`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}`
- **Success Response**: 200 OK
```json
{
    "users": [
        {
            "id": 1,
            "email": "ornek@email.com",
            "first_name": "Ad",
            "last_name": "Soyad",
            "role": "user",
            "is_active": true,
            "created_at": "2024-05-03T14:45:00Z",
            "updated_at": "2024-05-03T14:45:00Z"
        }
    ],
    "total": 1,
    "page": 1,
    "limit": 10
}
```

### Deactivate User (Admin Only)
- **URL**: `http://localhost:8080/users/{id}/deactivate`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
- **Success Response**: 200 OK

### Activate User (Admin Only)
- **URL**: `http://localhost:8080/users/{id}/activate`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
- **Success Response**: 200 OK

### Update User Role (Admin Only)
- **URL**: `http://localhost:8080/users/{id}/role`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "role": "admin"
}
```
- **Success Response**: 200 OK

### Reset User Password (Admin Only)
- **URL**: `http://localhost:8080/users/{id}/reset-password`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "new_password": "yeni456"
}
```
- **Success Response**: 200 OK

## Address Management Endpoints (All Protected)

### Create Address
- **URL**: `http://localhost:8080/users/addresses`
- **Method**: POST
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "type": "home",
    "title": "Ev Adresi",
    "address_line": "Örnek Mahallesi, Örnek Sokak No:1",
    "city": "İstanbul",
    "state": "Kadıköy",
    "country": "Türkiye",
    "postal_code": "34700",
    "is_default": true
}
```
- **Success Response**: 201 Created
```json
{
    "id": 1,
    "user_id": 1,
    "type": "home",
    "title": "Ev Adresi",
    "address_line": "Örnek Mahallesi, Örnek Sokak No:1",
    "city": "İstanbul",
    "state": "Kadıköy",
    "country": "Türkiye",
    "postal_code": "34700",
    "is_default": true,
    "created_at": "2024-05-03T14:45:00Z",
    "updated_at": "2024-05-03T14:45:00Z"
}
```

### List Addresses
- **URL**: `http://localhost:8080/users/addresses`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}`
- **Success Response**: 200 OK
```json
[
    {
        "id": 1,
        "user_id": 1,
        "type": "home",
        "title": "Ev Adresi",
        "address_line": "Örnek Mahallesi, Örnek Sokak No:1",
        "city": "İstanbul",
        "state": "Kadıköy",
        "country": "Türkiye",
        "postal_code": "34700",
        "is_default": true,
        "created_at": "2024-05-03T14:45:00Z",
        "updated_at": "2024-05-03T14:45:00Z"
    }
]
```

### Update Address
- **URL**: `http://localhost:8080/users/addresses/{id}`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "type": "work",
    "title": "İş Adresi",
    "address_line": "İş Merkezi, Kat:5",
    "city": "İstanbul",
    "state": "Levent",
    "country": "Türkiye",
    "postal_code": "34330",
    "is_default": true
}
```
- **Success Response**: 200 OK
```json
{
    "id": 1,
    "user_id": 1,
    "type": "work",
    "title": "İş Adresi",
    "address_line": "İş Merkezi, Kat:5",
    "city": "İstanbul",
    "state": "Levent",
    "country": "Türkiye",
    "postal_code": "34330",
    "is_default": true,
    "created_at": "2024-05-03T14:45:00Z",
    "updated_at": "2024-05-03T14:45:00Z"
}
```

### Delete Address
- **URL**: `http://localhost:8080/users/addresses/{id}`
- **Method**: DELETE
- **Headers**: 
  - `Authorization: Bearer {token}`
- **Success Response**: 204 No Content

## Contact Management Endpoints (All Protected)

### Create Contact
- **URL**: `http://localhost:8080/users/contacts`
- **Method**: POST
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "type": "home",
    "title": "Ev Telefonu",
    "phone_number": "+905551234567",
    "is_default": true
}
```
- **Success Response**: 201 Created
```json
{
    "id": 1,
    "user_id": 1,
    "type": "home",
    "title": "Ev Telefonu",
    "phone_number": "+905551234567",
    "is_default": true,
    "created_at": "2024-05-03T14:45:00Z",
    "updated_at": "2024-05-03T14:45:00Z"
}
```

### List Contacts
- **URL**: `http://localhost:8080/users/contacts`
- **Method**: GET
- **Headers**: 
  - `Authorization: Bearer {token}`
- **Success Response**: 200 OK
```json
[
    {
        "id": 1,
        "user_id": 1,
        "type": "home",
        "title": "Ev Telefonu",
        "phone_number": "+905551234567",
        "is_default": true,
        "created_at": "2024-05-03T14:45:00Z",
        "updated_at": "2024-05-03T14:45:00Z"
    }
]
```

### Update Contact
- **URL**: `http://localhost:8080/users/contacts/{id}`
- **Method**: PUT
- **Headers**: 
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "type": "work",
    "title": "İş Telefonu",
    "phone_number": "+905559876543",
    "is_default": true
}
```
- **Success Response**: 200 OK
```json
{
    "id": 1,
    "user_id": 1,
    "type": "work",
    "title": "İş Telefonu",
    "phone_number": "+905559876543",
    "is_default": true,
    "created_at": "2024-05-03T14:45:00Z",
    "updated_at": "2024-05-03T14:45:00Z"
}
```

### Delete Contact
- **URL**: `http://localhost:8080/users/contacts/{id}`
- **Method**: DELETE
- **Headers**: 
  - `Authorization: Bearer {token}`
- **Success Response**: 204 No Content

## Cart Endpoints (Tümü Korumalı)

### Get Cart
- **URL**: `http://localhost:8080/cart`
- **Method**: GET
- **Headers**:
  - `Authorization: Bearer {token}`
- **Success Response**: 200 OK
```json
{
    "id": 1,
    "user_id": 1,
    "items": [
        {
            "id": 1,
            "cart_id": 1,
            "product_id": 2,
            "quantity": 3,
            "created_at": "2024-05-03T14:45:00Z",
            "updated_at": "2024-05-03T14:45:00Z"
        }
    ],
    "created_at": "2024-05-03T14:45:00Z",
    "updated_at": "2024-05-03T14:45:00Z"
}
```

### Get Cart Items
- **URL**: `http://localhost:8080/cart/items`
- **Method**: GET
- **Headers**:
  - `Authorization: Bearer {token}`
- **Success Response**: 200 OK
```json
[
    {
        "id": 1,
        "cart_id": 1,
        "product_id": 2,
        "quantity": 3,
        "created_at": "2024-05-03T14:45:00Z",
        "updated_at": "2024-05-03T14:45:00Z"
    }
]
```

### Add Item to Cart
- **URL**: `http://localhost:8080/cart/items`
- **Method**: POST
- **Headers**:
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "product_id": 2,
    "quantity": 3
}
```
- **Success Response**: 201 Created

### Update Cart Item Quantity
- **URL**: `http://localhost:8080/cart/items/{id}`
- **Method**: PUT
- **Headers**:
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`
- **Body**:
```json
{
    "quantity": 5
}
```
- **Success Response**: 200 OK

### Remove Item from Cart
- **URL**: `http://localhost:8080/cart/items/{id}`
- **Method**: DELETE
- **Headers**:
  - `Authorization: Bearer {token}`
- **Success Response**: 204 No Content

### Clear Cart
- **URL**: `http://localhost:8080/cart`
- **Method**: DELETE
- **Headers**:
  - `Authorization: Bearer {token}`
- **Success Response**: 204 No Content

## Notes
1. Tüm protected endpoint'ler için `Authorization` header'ında geçerli bir JWT token gereklidir.
2. Token formatı: `Bearer {token}`
3. Başarılı login sonrası alınan token'ı diğer isteklerde kullanabilirsiniz.
4. Tüm endpoint'lerde hata durumunda uygun HTTP status code'ları ve hata mesajları döner. 
5. Adres ve iletişim bilgilerinde `type` alanı "home" veya "work" olabilir.
6. Adres ve iletişim bilgilerinde sadece bir tane varsayılan (default) kayıt olabilir.
7. Admin yetkisi gerektiren endpoint'ler için kullanıcının "admin" rolüne sahip olması gerekir. 