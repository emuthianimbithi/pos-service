# POS Service - Multi-Tenant REST API

A comprehensive Point of Sale (POS) system REST API built with **Gin** and **GORM**, featuring multi-tenant architecture at business and branch levels.

## 🏗️ Architecture

This is a **skeleton project** with a complete structure ready for implementation. The architecture follows clean architecture principles with clear separation of concerns:

```
pos-service/
├── cmd/api/                    # Application entry point
├── internal/
│   ├── config/                 # Configuration management
│   ├── database/               # Database connection & migrations
│   ├── middleware/             # HTTP middleware
│   │   ├── auth.go            # JWT authentication
│   │   ├── tenant.go          # Multi-tenant context
│   │   ├── ratelimit.go       # Rate limiting
│   │   ├── audit.go           # Request auditing
│   │   ├── logger.go          # Request logging
│   │   └── cors.go            # CORS & recovery
│   ├── models/                 # GORM models
│   │   ├── audit.go           # ✅ Audit log model (implemented)
│   │   └── models.go          # TODO: Domain models
│   ├── handlers/               # HTTP handlers (TODO)
│   ├── services/               # Business logic (TODO)
│   ├── repository/             # Data access layer (TODO)
│   └── utils/                  # Utility functions
│       ├── jwt.go             # JWT token utilities
│       ├── password.go        # Password hashing
│       └── pagination.go      # Pagination helpers
└── pkg/response/               # Standard API responses
```

## ✨ Features

### Implemented
- ✅ **Complete project structure** with all directories
- ✅ **Middleware stack**:
  - JWT authentication with role-based access control
  - Multi-tenant context (business & branch isolation)
  - Rate limiting (per-IP and per-user)
  - Request auditing (logs all API calls to database)
  - Request logging
  - CORS support
  - Panic recovery
- ✅ **Audit logging system** - tracks all requests with user context
- ✅ **Configuration management** - environment-based config
- ✅ **Database setup** - PostgreSQL with GORM
- ✅ **Utility functions** - JWT, password hashing, pagination
- ✅ **Standard API responses** - success, error, paginated responses
- ✅ **Router setup** - organized route groups with middleware

### To Implement
- 📝 Domain models (Business, Branch, User, Product, Sale, etc.)
- 📝 Repository layer (data access)
- 📝 Service layer (business logic)
- 📝 HTTP handlers (API endpoints)
- 📝 Input validation
- 📝 Unit tests

## 🚀 Getting Started

### Prerequisites
- Go 1.23 or higher
- PostgreSQL 12 or higher

### Installation

1. **Clone and navigate to the project**
   ```bash
   cd pos-service
   ```

2. **Install dependencies**
   ```bash
   make deps
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Create database**
   ```bash
   createdb pos_db
   ```

5. **Run the application**
   ```bash
   make run
   ```

The server will start on `http://localhost:8080`

## 🔧 Configuration

Edit `.env` file with your settings:

```env
# Server
ENVIRONMENT=development
PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=pos_db
DB_SSLMODE=disable

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24

# Rate Limiting
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=60
```

> **Note:** Environment variables are loaded automatically by the Makefile when using `make run` or `make build`.

## 🏢 Multi-Tenant Architecture

The system supports two levels of tenancy:

1. **Business Level** - Top-level tenant (e.g., a retail chain)
2. **Branch Level** - Sub-tenant (e.g., individual stores)

### Tenant Context

Tenant context is automatically extracted from:
- JWT claims (set during authentication)
- HTTP headers (`X-Business-ID`, `X-Branch-ID`)

### Middleware Usage

```go
// Require business context
router.Use(middleware.RequireBusiness())

// Require branch context
router.Use(middleware.RequireBranch())

// Require specific role
router.Use(middleware.RequireRole("admin", "manager"))
```

## 🔐 Authentication

JWT-based authentication with role support:

**Roles:**
- `admin` - Full system access
- `manager` - Branch management
- `cashier` - POS operations

**Token Claims:**
```json
{
  "user_id": "uuid",
  "business_id": "uuid",
  "branch_id": "uuid",
  "role": "admin"
}
```

## 📊 Audit Logging

All API requests are automatically logged to the `audit_logs` table with:
- Request details (method, path, status, IP)
- User context (user_id, business_id, branch_id)
- Performance metrics (duration)
- Action metadata (action, resource, resource_id)

### Setting Audit Context

```go
middleware.SetAuditAction(c, "create", "product", productID.String())
```

## 🛣️ API Routes (Planned)

```
GET  /health                    # Health check

POST /api/v1/auth/register      # User registration
POST /api/v1/auth/login         # User login
POST /api/v1/auth/refresh       # Refresh token

# Business routes (admin only)
GET    /api/v1/businesses       # List businesses
POST   /api/v1/businesses       # Create business
GET    /api/v1/businesses/:id   # Get business
PUT    /api/v1/businesses/:id   # Update business
DELETE /api/v1/businesses/:id   # Delete business

# Branch routes (business scoped)
GET    /api/v1/branches         # List branches
POST   /api/v1/branches         # Create branch
GET    /api/v1/branches/:id     # Get branch
PUT    /api/v1/branches/:id     # Update branch
DELETE /api/v1/branches/:id     # Delete branch

# Product routes (branch scoped)
GET    /api/v1/products         # List products
POST   /api/v1/products         # Create product
GET    /api/v1/products/:id     # Get product
PUT    /api/v1/products/:id     # Update product
DELETE /api/v1/products/:id     # Delete product

# Sales routes (branch scoped)
GET    /api/v1/sales            # List sales
POST   /api/v1/sales            # Create sale
GET    /api/v1/sales/:id        # Get sale

# Customer routes (branch scoped)
GET    /api/v1/customers        # List customers
POST   /api/v1/customers        # Create customer
GET    /api/v1/customers/:id    # Get customer
PUT    /api/v1/customers/:id    # Update customer
DELETE /api/v1/customers/:id    # Delete customer
```

## 🔨 Development

### Available Make Commands

```bash
make help          # Show all available commands
make build         # Build the application
make run           # Run the application
make dev           # Run with hot reload (requires air)
make test          # Run tests
make test-coverage # Run tests with coverage report
make clean         # Clean build artifacts
make deps          # Download dependencies
make lint          # Run linter (requires golangci-lint)
make fmt           # Format code
make vet           # Run go vet
```

### Build
```bash
make build
```

### Run
```bash
make run
```

### Test (when implemented)
```bash
make test
```

## 📝 Next Steps

1. **Define Models** - Implement domain models in `internal/models/models.go`
2. **Create Repositories** - Implement data access in `internal/repository/`
3. **Build Services** - Implement business logic in `internal/services/`
4. **Add Handlers** - Implement HTTP handlers in `internal/handlers/`
5. **Uncomment Routes** - Enable routes in `internal/handlers/router.go`
6. **Add Validation** - Implement input validation
7. **Write Tests** - Add unit and integration tests

## 📦 Dependencies

- **[Gin](https://github.com/gin-gonic/gin)** - HTTP web framework
- **[GORM](https://gorm.io/)** - ORM library
- **[JWT](https://github.com/golang-jwt/jwt)** - JSON Web Tokens
- **[UUID](https://github.com/google/uuid)** - UUID generation
- **[Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** - Password hashing
- **[Rate](https://pkg.go.dev/golang.org/x/time/rate)** - Rate limiting

> Environment variables are loaded via Makefile, no additional packages needed.

## 📄 License

MIT License
