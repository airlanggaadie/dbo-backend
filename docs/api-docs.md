# API DOCUMENTATION

- [API DOCUMENTATION](#api-documentation)
  - [Login User](#login-user)
  - [Get Login Report](#get-login-report)
  - [Get List of User (paginated)](#get-list-of-user-paginated)
  - [Get User Detail](#get-user-detail)
  - [Search User by Name](#search-user-by-name)
  - [Add New User](#add-new-user)
  - [Update User](#update-user)
  - [Delete User](#delete-user)
  - [Get List of Order (paginated)](#get-list-of-order-paginated)
  - [Get Order Detail](#get-order-detail)
  - [Search Order by Order Number](#search-order-by-order-number)
  - [Add New Order](#add-new-order)
  - [Update Order](#update-order)
  - [Delete Order](#delete-order)

## Login User

Authentication user before access the management systen

method: `POST`

path: `/auth/login`

header: `Content-Type: application/json`

body: `json`
```jsonc
{
    "username": "admin",        // string, required
    "password": "supersecret"   // string, required
}
```

response: `json`
```jsonc
// status code 200
{
    "token": "supertoken"
}

// status code 400
{
    "message": "invalid request"
}

// status code 500
{
    "message": "internal server error"
}
```

sample:
```bash
curl --request POST --url 'http://localhost:3000/auth/login' \
    --header 'Content-Type: application/json' \
    --data '{"username":"admin","password":"admin-password"}'
```

## Get Login Report

Get login history for all users that are logged in

method: `GET`

path: `/auth/login/report`

headers: 
- Content-Type: application/json
- Authorization: Bearer supertoken

query parameters:
- page: page number, optional, numeric, default value is 1
- limit: size per page, optional, numeric, default value is 10

response:
```jsonc
// status code 200
{
    "data": [
        {
            "username": "admin",
            "login_time": "2024-03-03 00:00:00Z"
        }
    ]
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}
```

sample:
```bash
curl --request POST --url 'http://localhost:3000/auth/login/report?page=1&limit=10'
```

## Get List of User (paginated)

Get list of users that paginated

method: `GET`

path: `/users`

headers:
- Authorization: Bearer supertoken

query parameters:
- page: page number, optional, numeric, default value is 1
- limit: size per page, optional, numberic, default value is 10
- q: filter by name, optional, alphabetic (insensitive case)

response:
```jsonc
// status code 200
{
    "data": [
        {
            "id": "uuid-admin",
            "username": "admin",
            "name": "Admin",
            "created_at": "2024-03-03 00:00:00Z",
            "updated_at": "2024-03-03 00:00:00Z"
        }
    ]
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}
```

sample:
```bash
curl --request POST --url 'http://localhost:3000/users?page=1&limit=10&q=admin'
```

## Get User Detail

Get user information

method: `GET`

path: `/user/:id`

headers:
- Authorization: Bearer supertoken

response:
```jsonc
// status code 200
{
    "id": "uuid-admin",
    "username": "admin",
    "name": "Admin",
    "created_at": "2024-03-03 00:00:00Z",
    "updated_at": "2024-03-03 00:00:00Z"
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}

// status code 404
{
    "message": "not found"
}
```

sample:
```bash
curl --request POST --url 'http://localhost:3000/user/uuid-admin'
```

## Search User by Name

simple search user by name

method: `GET`

path: `/user/search`

headers:
- Authorization: Bearer supertoken

query parameters:
- q: filter by name, optional, alphabetic (insensitive case)

response:
```jsonc
// status code 200
{
    "data": [
        {
            "id": "uuid-admin",
            "name": "Admin",
        }
    ]
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}
```

sample:
```bash
curl --request POST --url 'http://localhost:3000/user/search?q=admin'
```

## Add New User

create new user with default password

method: `POST`

path: `/user`

headers:
- Content-Type: application/json
- Authorization: Bearer supertoken

body: `json`
```jsonc
{
    "username": "newusername", // unique username, got 400 when already registered
    "name": "newuser",
}
```

response:
```jsonc
// status code 200
{
    "id": "uuid-newuser",
    "username": "newusername",
    "name": "newuser",
    "created_at": "2024-03-03 00:00:00Z",
    "updated_at": "2024-03-03 00:00:00Z"
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}
```

sample:
```bash
curl --request POST --url 'http://localhost:3000/user' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer supertoken' \
--data '{"username":"newusername","name":"newuser"}'
```

## Update User

update user detail

method: `PUT`

path: `/user/:id`

headers:
- Content-Type: application/json
- Authorization: Bearer supertoken

body: `json`
```jsonc
{
    "username": "newusername", // unique username, got 400 when already registered
    "name": "newuserupdate",
}
```

response:
```jsonc
// status code 200
{
    "id": "uuid-newuser",
    "username": "newusername",
    "name": "newuserupdate",
    "created_at": "2024-03-03 00:00:00Z",
    "updated_at": "2024-03-03 00:00:00Z"
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}

// status code 404
{
    "message": "not found"
}
```

sample:
```bash
curl --request PUT --url 'http://localhost:3000/user/uuid-newuser' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer supertoken' \
--data '{"username":"newusername","name":"newuserupdate"}'
```

## Delete User

deleting user data by user id

method: `DELETE`

path: `/order/:id`

headers:
- Authorization: Bearer supertoken

response:
```jsonc
// status code 204 No Content

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}

// status code 404
{
    "message": "not found"
}
```

sample:
```bash
curl --request PUT --url 'http://localhost:3000/user/uuid-newuser' \
--header 'Authorization: Bearer supertoken'
```

## Get List of Order (paginated)

Get list of orders that paginated

method: `GET`

path: `/orders`

headers:
- Authorization: Bearer supertoken

query parameters:
- page: page number, optional, numeric, default value is 1
- limit: size per page, optional, numberic, default value is 10
- q: filter by order number, optional, alphabetic (insensitive case)

response:
```jsonc
// status code 200
{
    "data": [
        {
            "id": "uuid-order-id-1",
            "buyer_name": "buyer",
            "code": "order-short-number",
            "total_item": 2,
            "total_price": 600000,
            "created_at": "2024-03-03T00:00:00Z",
            "updated_at": "2024-03-03T00:00:00Z"
        }
    ]
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}
```

sample:
```bash
curl --request POST --url 'http://localhost:3000/orders?page=1&limit=10&q=admin'
```

## Get Order Detail

Get order summary information

method: `GET`

path: `/order/:id`

headers:
- Authorization: Bearer supertoken

response:
```jsonc
// status code 200
{
    "id": "uuid-order-id-1",
    "buyer_name": "buyer",
    "code": "order-short-number",
    "items": [
        {
            "id": "uuid-item-1",
            "code": "k-001",
            "name": "kursi",
            "quantity": 2,
            "unit_price": 300000,
            "sub_total": 600000,
            "created_at": "2024-03-03T00:00:00Z",
            "updated_at": "2024-03-03T00:00:00Z"
        }
    ],
    "total_item": 2,
    "total_price": 600000,
    "created_at": "2024-03-03T00:00:00Z",
    "updated_at": "2024-03-03T00:00:00Z"
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}

// status code 404
{
    "message": "not found"
}
```

sample:
```bash
curl --request POST --url 'http://localhost:3000/order/uuid-order-1'
```

## Search Order by Order Number

simple search order by order number

method: `GET`

path: `/order/search`

headers:
- Authorization: Bearer supertoken

query parameters:
- q: filter by order short number, optional, alphabetic (insensitive case)

response:
```jsonc
// status code 200
{
    "data": [
        {
            "id": "uuid-order-1",
            "code": "order-short-number",
        }
    ]
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}
```

sample:
```bash
curl --request POST --url 'http://localhost:3000/order/search?q=short'
```

## Add New Order

add new order with buyer name and detail items

method: `POST`

path: `/order`

headers:
- Content-Type: application/json
- Authorization: Bearer supertoken

body: `json`
```jsonc
{
    "buyer_name": "buyer",      // required, string
    "items": [
        {
            "code": "k-001",    // required, string
            "name": "kursi",    // required, string
            "quantity": 2,      // required, numeric
            "price": 200000     // unit price, required, numbric
        }
    ]
}
```

response:
```jsonc
// status code 200
{
    "id": "uuid-order-id-1",
    "buyer_name": "buyer",
    "code": "order-short-number",
    "items": [
        {
            "id": "uuid-item-1",
            "code": "k-001",
            "name": "kursi",
            "quantity": 2,
            "unit_price": 200000,
            "sub_total": 40000,
        }
    ],
    "total_item": 2,
    "total_price": 40000
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}
```

sample:
```bash
curl --request POST --url 'http://localhost:3000/order' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer supertoken' \
--data '{"buyer_name":"buyer","item":[{"code":"k-001","name":"kursi","quantity":2,"unit_price":200000}]}'
```

## Update Order

update buyer name and detail of items (add or remove), when updated items is empty, it will delete the order.

method: `PUT`

path: `/order/:id`

headers:
- Content-Type: application/json
- Authorization: Bearer supertoken

body: `json`
```jsonc
{
    "buyer_name": "buyer",          // required, string
}
```

response:
```jsonc
// status code 200
{
    "id": "uuid-order-id-1",
    "buyer_name": "buyer",
    "code": "order-short-number",
    "items": [
        {
            "id": "uuid-item-1",
            "code": "k-001",
            "name": "kursi",
            "quantity": 2,
            "unit_price": 300000,
            "sub_total": 600000,
        }
    ],
    "total_item": 2,
    "total_price": 600000
}

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}

// status code 404
{
    "message": "not found"
}
```

sample:
```bash
curl --request PUT --url 'http://localhost:3000/order/uuid-order-id-1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer supertoken' \
--data '{"id":"uuid-order-id-1","buyer_name":"buyer","item":[{"id": "uuid-item-1","code":"k-001","name":"kursi","quantity":2,"unit_price":200000}]}'
```

## Delete Order

deleting order data by order id

method: `DELETE`

path: `/order/:id`

headers:
- Authorization: Bearer supertoken

response:
```jsonc
// status code 204 No Content

// status code 400
{
    "message": "invalid request"
}

// status code 401
{
    "message": "unauthorize"
}

// status code 404
{
    "message": "not found"
}
```

sample:
```bash
curl --request PUT --url 'http://localhost:3000/order/uuid-order-1' \
--header 'Authorization: Bearer supertoken'
```
