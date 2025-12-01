# Quick Start Guide

## Setup

1. **Copy environment file**
   ```bash
   cp .env.example .env
   ```

2. **Edit `.env` with your database credentials**
   ```bash
   # Update these values in .env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=pos_db
   ```

3. **Install dependencies**
   ```bash
   make deps
   ```

4. **Create database**
   ```bash
   createdb pos_db
   ```

## Running the Application

### Development
```bash
make run
```
The server will start on `http://localhost:8080`

### Build and Run
```bash
make build
./bin/api
```

## Available Commands

Run `make help` to see all available commands:

```bash
make help          # Show all available commands
make build         # Build the application
make run           # Run the application (loads .env automatically)
make dev           # Run with hot reload (requires air)
make test          # Run tests
make test-coverage # Run tests with coverage report
make clean         # Clean build artifacts
make deps          # Download dependencies
make lint          # Run linter (requires golangci-lint)
make fmt           # Format code
make vet           # Run go vet
```

## Testing the API

Once running, test the health endpoint:
```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "healthy",
  "service": "pos-service"
}
```

## Environment Variables

The Makefile automatically loads environment variables from `.env` when you run:
- `make run`
- `make build`
- Any other make commands

No need for additional packages like godotenv!

## Next Steps

1. Implement your domain models in `internal/models/models.go`
2. Create repositories in `internal/repository/`
3. Build services in `internal/services/`
4. Add handlers in `internal/handlers/`
5. Uncomment routes in `internal/handlers/router.go`

See [README.md](README.md) for complete documentation.
