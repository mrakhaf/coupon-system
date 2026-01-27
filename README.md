# Coupon System API

A clean architecture implementation for a coupon management system using Go, Echo framework, and MySQL database.

## Features

- **Clean Architecture**: Separated into distinct layers (Domain, Use Cases, Infrastructure, Presentation)
- **RESTful API**: Complete CRUD operations for coupon management
- **Coupon Usage**: Apply coupons to purchases with validation
- **Database**: MySQL with proper indexing and constraints
- **Validation**: Input validation using struct tags
- **Error Handling**: Comprehensive error handling throughout the application

## Project Structure

```
coupon-system/
├── cmd/
│   └── api/                    # Main application entry point
├── internal/
│   ├── config/                 # Configuration management
│   ├── domain/                 # Domain entities and interfaces
│   │   ├── coupon.go          # Coupon entity and request/response structs
│   │   ├── repository.go      # Repository interface
│   │   └── usecase.go         # Use case interface
│   ├── infrastructure/         # Database and external services
│   │   ├── database.go        # Database connection
│   │   └── coupon_repository.go # Repository implementation
│   ├── usecase/               # Business logic
│   │   └── coupon_usecase.go  # Use case implementation
│   └── handler/               # HTTP handlers
│       └── coupon_handler.go  # API endpoints
├── db/
│   └── migrations/            # Database migration scripts
└── go.mod                    # Go module definition
```

## API Endpoints

### Coupons

- `POST /api/v1/coupons` - Create a new coupon
- `GET /api/v1/coupons/:id` - Get coupon by ID
- `GET /api/v1/coupons/code/:code` - Get coupon by code
- `PUT /api/v1/coupons/:id` - Update coupon
- `DELETE /api/v1/coupons/:id` - Delete coupon
- `GET /api/v1/coupons` - List coupons with pagination
- `POST /api/v1/coupons/use` - Apply coupon to purchase

## Installation

### Prerequisites

- Go 1.23+
- MySQL database

### Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd coupon-system
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your database configuration
   ```

4. Run database migrations:
   ```bash
   # Connect to your MySQL database and run the SQL from db/migrations/
   mysql -u username -p database_name < db/migrations/001_create_coupons_table.sql
   ```

5. Start the application:
   ```bash
   go run cmd/api/main.go
   ```

## Configuration

The application uses environment variables for configuration. Copy `.env.example` to `.env` and modify the values:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password_here
DB_NAME=coupon_system

# Application Configuration
PORT=8080
```

## Usage Examples

### Create a Coupon

```bash
curl -X POST http://localhost:8080/api/v1/coupons \
  -H "Content-Type: application/json" \
  -d '{
    "code": "SAVE20",
    "description": "20% off on all products",
    "discount": 20.0,
    "min_amount": 100.0,
    "max_usage": 50
  }'
```

### Apply a Coupon

```bash
curl -X POST http://localhost:8080/api/v1/coupons/use \
  -H "Content-Type: application/json" \
  -d '{
    "code": "SAVE20",
    "amount": 150.0
  }'
```

## Testing

To run tests:

```bash
go test ./...
```

## Architecture Overview

### Domain Layer
Contains business entities, value objects, and domain interfaces. This layer is independent of any framework or database.

### Use Case Layer
Contains business logic and orchestrates the flow between domain entities and infrastructure.

### Infrastructure Layer
Contains database implementations, external service integrations, and other technical details.

### Presentation Layer
Contains HTTP handlers and API endpoints that interact with the use cases.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for your changes
5. Submit a pull request

## License

This project is licensed under the MIT License.