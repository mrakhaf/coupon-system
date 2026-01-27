# Coupon System API

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

The application uses environment variables for configuration. Copy `.env.example` to `.env` 
