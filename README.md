# Coupon System

A Go-based coupon management system with REST API endpoints for creating and claiming coupons.

## Features

- Create coupons with unique names and amounts
- Claim coupons by name with user identification
- Database-backed with MySQL
- RESTful API design
- Input validation

## Tech Stack

- Go 1.24
- Echo framework
- GORM ORM
- MySQL
- Docker & Docker Compose

## Quick Start

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- MySQL (optional, can use Docker)

### Using Docker (Recommended)

1. Clone the repository
2. Build and start services:

```bash
docker compose up -d
```

3. The API will be available at `http://localhost:8080`
4. Database will be available at `http://localhost:3306`
5. Adminer (database management) at `http://localhost:8081`

### Manual Setup

1. Clone the repository
2. Install dependencies:

```bash
go mod download
```

3. Set up environment variables (copy `.env.example` to `.env` and configure)
4. Run the application:

```bash
go run cmd/api/main.go
```

## API Endpoints

### Health Check

```http
GET /
```

Returns system status and version.

### Create Coupon

```http
POST /api/coupons
Content-Type: application/json

{
    "name": "PROMO_1",
    "amount": 100.00
}
```

**Response:**
- `201 Created` - Coupon created successfully
- `400 Bad Request` - Invalid input
- `409 Conflict` - Coupon name already exists

### Claim Coupon

```http
POST /api/coupons/claim
Content-Type: application/json

{
    "name": "PROMO_1",
    "user_id": "user123"
}
```

**Response:**
- `200 OK` - Coupon claimed successfully
- `400 Bad Request` - Invalid input
- `404 Not Found` - Coupon not found or inactive
- `409 Conflict` - Coupon already claimed by user

### Get Coupon Details

```http
GET /api/coupons/{name}
```

**Response:**
- `200 OK` - Coupon details
- `404 Not Found` - Coupon not found

## Database Schema

The system uses two main tables:

### coupons
- `id` - Unique identifier (UUID)
- `name` - Coupon name (unique)
- `amount` - Coupon value
- `is_active` - Whether coupon can be claimed
- `created_at` - Creation timestamp
- `updated_at` - Last update timestamp

### coupon_claims
- `id` - Unique identifier (UUID)
- `user_id` - User who claimed the coupon
- `coupon_id` - Reference to claimed coupon
- `created_at` - Claim timestamp

## Environment Variables

- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 3306)
- `DB_USER` - Database username
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `PORT` - API server port (default: 8080)

## Docker Configuration

The project includes Docker support with:

- Multi-stage Dockerfile for optimized builds
- Docker Compose for easy local development
- MySQL database with automatic migration
- Adminer for database management

## Development

### Running Tests

```bash
go test ./...
```

### Code Formatting

```bash
go fmt ./...
```

### Building

```bash
go build -o coupon-system ./cmd/api
```

## License

MIT License


## Stress Test Using k6
https://github.com/mrakhaf/coupon-system-k6 