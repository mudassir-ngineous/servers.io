# Go REST API Server

A modern REST API server built with Go and Gin framework.

## Features

- RESTful API endpoints
- Middleware support (logging, recovery, CORS)
- Configuration management
- API versioning
- Health check endpoint
- User management endpoints (CRUD operations)

## Prerequisites

- Go 1.21 or higher
- Git

## Getting Started

1. Clone the repository:
```bash
git clone <repository-url>
cd web-server
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start on port 8080 by default. You can change the port by setting the `PORT` environment variable.

## API Endpoints

### Health Check
- `GET /api/v1/health` - Check if the service is healthy

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get a specific user
- `POST /api/v1/users` - Create a new user
- `PUT /api/v1/users/:id` - Update a user
- `DELETE /api/v1/users/:id` - Delete a user

## Environment Variables

- `PORT` - Server port (default: 8080)
- `LOG_LEVEL` - Logging level (default: info)

## Project Structure

```
web-server/
├── api/
│   ├── handlers/
│   │   └── handlers.go
│   └── router.go
├── config/
│   └── config.go
├── middleware/
│   └── middleware.go
├── models/
├── services/
├── main.go
└── README.md
```

## TODO

- [ ] Add database integration
- [ ] Implement user authentication
- [ ] Add request validation
- [ ] Add API documentation
- [ ] Add tests
- [ ] Add Docker support 