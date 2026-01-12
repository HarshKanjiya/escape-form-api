# EscapeForm API

A scalable Golang REST API built with Fiber framework, featuring JWT authentication, role-based authorization, and PostgreSQL database integration.

## Features

- **Framework**: Fiber v2 (high-performance Go web framework)
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT-based authentication
- **Authorization**: Role-based access control (RBAC)
- **Validation**: Request validation with struct tags
- **Logging**: Structured logging with Zerolog
- **Rate Limiting**: IP-based rate limiting
- **CORS**: Configurable CORS settings
- **Environment Management**: Multi-environment configuration (.env files)
- **Error Handling**: Centralized error handling with custom responses
- **API Documentation**: Auto-generated Swagger/OpenAPI documentation

## Project Structure

```
.
├── cmd/
│   └── app/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go               # Configuration management
│   ├── database/
│   │   └── postgres.go             # Database connection and migrations
│   ├── handlers/
│   │   └── health.go               # Health check endpoint
│   ├── middlewares/
│   │   ├── auth.go                 # JWT authentication middleware
│   │   ├── logger.go               # Logging middleware
│   │   ├── ratelimit.go            # Rate limiting middleware
│   │   ├── rbac.go                 # Role-based authorization
│   │   └── setup.go                # Middleware setup
│   ├── models/
│   │   ├── user.go                 # User model and DTOs
│   │   └── common.go               # Common models and utilities
│   ├── controllers/
│   │   └── user_controller.go      # User API endpoints
│   ├── routes/
│   │   └── routes.go               # Route definitions
│   └── services/
│       └── user_service.go         # Business logic layer
├── pkg/
│   └── utils/
│       ├── response.go             # HTTP response helpers
│       ├── errors.go               # Error handling utilities
│       ├── validator.go            # Request validation
│       ├── password.go             # Password hashing
│       └── jwt.go                  # JWT token utilities
├── configs/                        # Configuration files
├── scripts/                        # Utility scripts
├── tests/                          # Test files
├── .env.local                      # Local environment variables
├── .env.dev                        # Development environment variables
├── .env.prod                       # Production environment variables
├── .env.example                    # Environment variables template
└── go.mod                          # Go module file
```

## Prerequisites

- Go 1.19 or higher
- PostgreSQL database
- Git

## Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd escape-form-api
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Set up environment variables**

   ```bash
   cp .env.example .env.local
   # Edit .env.local with your database credentials and other settings
   ```

4. **Set up PostgreSQL database**

   - Create a PostgreSQL database
   - Update the database configuration in `.env.local`

5. **Run the application**

   ```bash
   go run cmd/app/main.go
   ```

   The server will start on the port specified in your environment configuration (default: 3000).

   **🚀 Local Development Bonus:** When running in `local` environment, the Swagger UI will automatically open in your default browser!

## Environment Configuration

The application supports multiple environments:

- **local**: For local development
- **dev**: For development environment
- **prod**: For production environment

Set the `APP_ENV` variable to switch between environments:

```bash
export APP_ENV=local  # or dev, prod
```

### Environment Variables

| Variable                | Description                  | Default                                                                                                                                                            |
| ----------------------- | ---------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `APP_ENV`               | Environment (local/dev/prod) | local                                                                                                                                                              |
| `APP_NAME`              | Application name             | EscapeForm API                                                                                                                                                     |
| `APP_PORT`              | Server port                  | 3000                                                                                                                                                               |
| `APP_HOST`              | Server host                  | localhost                                                                                                                                                          |
| `DB_HOST`               | Database host                | localhost                                                                                                                                                          |
| `DB_PORT`               | Database port                | 5432                                                                                                                                                               |
| `DB_USER`               | Database user                | postgres                                                                                                                                                           |
| `DB_PASSWORD`           | Database password            |                                                                                                                                                                    |
| `DB_NAME`               | Database name                | escape_form                                                                                                                                                        |
| `DB_SSLMODE`            | SSL mode                     | disable                                                                                                                                                            |
| `JWT_SECRET`            | JWT signing secret           |                                                                                                                                                                    |
| `JWT_EXPIRY`            | JWT token expiry             | 24h                                                                                                                                                                |
| `JWT_REFRESH_EXPIRY`    | JWT refresh token expiry     | 168h                                                                                                                                                               |
| `CORS_ORIGINS`          | Allowed CORS origins         | http://localhost:3000,http://localhost:8080,http://127.0.0.1:3000,http://127.0.0.1:8080,https://escform.com,https://dashboard.escform.com,https://form.escform.com |
| `CORS_METHODS`          | Allowed CORS methods         | GET,POST,PUT,DELETE,OPTIONS                                                                                                                                        |
| `CORS_HEADERS`          | Allowed CORS headers         | Content-Type,Authorization                                                                                                                                         |
| `RATE_LIMIT_MAX`        | Rate limit max requests      | 100                                                                                                                                                                |
| `RATE_LIMIT_EXPIRATION` | Rate limit window (seconds)  | 60                                                                                                                                                                 |
| `LOG_LEVEL`             | Logging level                | info                                                                                                                                                               |

## API Endpoints

### Health Check

- **GET** `/health` - Check API health status

### Authentication

- **POST** `/api/v1/auth/register` - Register a new user
- **POST** `/api/v1/auth/login` - User login

### Users (Protected)

All user endpoints require JWT authentication.

- **GET** `/api/v1/users` - Get all users (paginated)
- **GET** `/api/v1/users/:id` - Get user by ID
- **PUT** `/api/v1/users/:id` - Update user (admin only)
- **DELETE** `/api/v1/users/:id` - Delete user (admin only)
- **GET** `/api/v1/users/search` - Search users

## API Documentation

The API includes auto-generated Swagger/OpenAPI documentation that you can access at:

- **Swagger UI**: `http://localhost:3000/swagger/index.html`
- **ReDoc**: `http://localhost:3000/swagger/doc.json`

### Generating Documentation

To regenerate the Swagger documentation after making changes to API annotations:

```bash
swag init -g cmd/app/main.go -o docs
```

## Authentication

The API uses JWT (JSON Web Tokens) for authentication:

1. **Register/Login** to get access and refresh tokens
2. **Include the access token** in the Authorization header:
   ```
   Authorization: Bearer <your-jwt-token>
   ```

## User Roles

- **user**: Basic user permissions
- **admin**: Full access to all endpoints

## Development

### Running Tests

```bash
go test ./...
```

### Building for Production

```bash
go build -o bin/app cmd/app/main.go
```

### Code Formatting

```bash
go fmt ./...
```

### Linting

```bash
go vet ./...
```

## Deployment

1. Set `APP_ENV=prod` in your production environment
2. Configure production database credentials
3. Set a strong `JWT_SECRET`
4. Use a reverse proxy (nginx) for production deployment
5. Enable SSL/TLS

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License.
