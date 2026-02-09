# Category API

A simple REST API for managing product categories built with Go.

## Getting Started

### Prerequisites

- Go 1.25.6 or higher

### Installation

1. Clone the repository
2. Navigate to the project directory
3. Run the server:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Available Endpoints

### Health Check

- **Method:** `GET`
- **Path:** `/health`
- **Description:** Check if the API is running
- **Response:**

```json
{
  "status": "OK",
  "message": "API category is running"
}
```

### Get All Categories

- **Method:** `GET`
- **Path:** `/api/categories`
- **Description:** Retrieve all categories
- **Query Parameters:**
  - `search` (optional): Filter categories by name or description (case-insensitive)
- **Examples:**
  - Get all categories: `GET /api/categories`
  - Search for categories: `GET /api/categories?search=electronics`
- **Response:**

```json
[
  {
    "id": 1,
    "name": "Electronics",
    "description": "Devices and gadgets"
  },
  {
    "id": 2,
    "name": "Books",
    "description": "Printed and digital books"
  },
  {
    "id": 3,
    "name": "Clothing",
    "description": "Apparel and accessories"
  }
]
```

### Create a New Category

- **Method:** `POST`
- **Path:** `/api/categories`
- **Description:** Create a new category
- **Request Body:**

```json
{
  "name": "Home & Garden",
  "description": "Furniture and household items"
}
```

- **Response:** `201 Created`

```json
{
  "id": 4,
  "name": "Home & Garden",
  "description": "Furniture and household items"
}
```

### Get a Specific Category

- **Method:** `GET`
- **Path:** `/api/categories/{id}`
- **Description:** Retrieve a category by ID
- **Example:** `GET /api/categories/1`
- **Response:**

```json
{
  "id": 1,
  "name": "Electronics",
  "description": "Devices and gadgets"
}
```

### Update a Category

- **Method:** `PUT`
- **Path:** `/api/categories/{id}`
- **Description:** Update an existing category
- **Example:** `PUT /api/categories/1`
- **Request Body:**

```json
{
  "name": "Electronics & Gadgets",
  "description": "Updated description"
}
```

- **Response:**

```json
{
  "id": 1,
  "name": "Electronics & Gadgets",
  "description": "Updated description"
}
```

### Delete a Category

- **Method:** `DELETE`
- **Path:** `/api/categories/{id}`
- **Description:** Delete a category by ID
- **Example:** `DELETE /api/categories/1`
- **Response:**

```json
{
  "message": "Category deleted successfully"
}
```

## Data Model

```go
type Category struct {
  ID          int    `json:"id"`
  Name        string `json:"name"`
  Description string `json:"description"`
}
```

## Error Handling

- **400 Bad Request:** Invalid input or malformed request
- **404 Not Found:** Category not found
- **201 Created:** Successfully created a new category

## Default Categories

The API comes with three pre-loaded categories:

1. Electronics - Devices and gadgets
2. Books - Printed and digital books
3. Clothing - Apparel and accessories
