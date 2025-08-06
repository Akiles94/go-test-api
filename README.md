# 🚀 Go Test API

A modern demo REST API built with Go, Gin, GORM, and PostgreSQL following Clean Architecture and Domain-Driven Design principles.

## 📋 Features

- ✅ **Clean Architecture** - Clear separation of concerns
- ✅ **Complete REST API** - Product CRUD operations with pagination
- ✅ **Swagger Documentation** - Self-documented API
- ✅ **Dockerized** - Simplified development and deployment
- ✅ **Helm Chart** - Kubernetes-ready
- ✅ **Middlewares** - Rate limiting, CORS, structured logging
- ✅ **Health checks** - Monitoring and observability

## 🛠️ Tech Stack

- **Framework:** Gin (Go)
- **ORM:** GORM
- **Database:** PostgreSQL
- **Documentation:** Swagger/OpenAPI
- **Containers:** Docker + Docker Compose
- **Orchestration:** Kubernetes + Helm

## 🚀 Quick Start

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- Make (optional)

### 1. Clone the repository

```bash
git clone https://github.com/Akiles94/go-test-api.git
cd go-test-api
```

### 2. Set up env variables

cp .env.example .env

# Edit .env with your configurations

### 3. Run with Docker Compose

# Start the complete application stack

docker-compose up -d

# View logs

docker-compose logs -f go-test-api

### 4. Verify it's working

- API Health: http://localhost:8080/health
- Swagger UI: http://localhost:8080/swagger/index.html
- API Endpoint: http://localhost:8080/api/v1/products

### Project Structure

go-test-api/
├── cmd/ # Application entry point
│ └── main.go
├── contexts/ # Domain contexts
│ ├── products/
│ │ ├── application/ # Use cases
│ │ │ └── use_cases/
│ │ ├── domain/ # Entities and repositories
│ │ └── infra/ # Implementations
│ │ ├── adapters/ # Repositories
│ │ ├── handlers/ # HTTP controllers
│ │ └── modules/ # Module configuration
│ └── shared/ # Shared code
│ ├── application/
│ └── infra/
│ └── middlewares/
├── config/ # Configuration
├── db/ # Database connection
├── docs/ # Generated Swagger documentation
├── helm-chart/ # Helm chart for Kubernetes
├── k8s/ # Kubernetes manifests
├── Dockerfile # Production image
├── docker-compose.yaml # Development environment
└── README.md

### Local development

# Install dependencies

go mod download

# Generate Swagger documentation

swag init -g cmd/main.go -o ./docs

# Run the application with hot reload

make dev

### 🧪 Testing

make test

### 🧪 Health check

curl http://localhost:8080/health

### Basic metrics

- Structured logging
- Rate limiting
- Request ID tracking
- Recovery middleware

### Security

✅ Non-privileged user in container
✅ Scratch-based image (minimal attack surface)
✅ Externalized secrets
✅ Configured rate limiting
✅ Security headers
