# Mogi API

Mogi adalah aplikasi untuk mock response API yang dibangun dengan Go mengikuti pola Clean Architecture. Aplikasi ini memudahkan developer untuk membuat mock endpoint dengan response yang dapat dikustomisasi.

## Struktur Project

```
Mogi/
├── main.go                                   # Application entry point
├── internal/
│   ├── domain/
│   │   └── user.go                         # User domain model
│   ├── dto/
│   │   └── users.go                       # User DTOs
│   ├── infrastructure/
│   │   └── postgres/
│   │       └── postgres.go                 # Database connection & repository
│   ├── interfaces/
│   │   ├── handler/
│   │   │   └── users.go                  # User HTTP handlers
│   │   └── server/
│   │       └── server.go                  # Echo server configuration
│   ├── shared/
│   │   ├── config/
│   │   │   └── config.go                  # Configuration loader
│   │   ├── error/
│   │   │   └── error.go                   # Error handling
│   │   ├── constants/                     # Constants
│   │   ├── utils/                         # Utility functions
│   │   └── validator/                     # Validation
│   └── usecase/
│       └── users/
│           └── usecase.go                 # User business logic
├── migrations/                              # Database migrations
│   └── 001_create_users_table.sql          # Users table migration
├── resources/                              # Configuration & resource files
│   ├── config.json                        # Configuration file
│   └── config.json.example                # Example configuration
├── go.mod                                 # Go module file
├── go.sum                                 # Go dependencies
├── Makefile                               # Build commands
└── README.md                              # This file
```

## Architecture Pattern

Project ini mengikuti **Clean Architecture** dengan separation of concerns:

### Layers:

1. **Domain Layer** (`internal/domain/`)
   - Berisi business entities dan rules
   - Tidak memiliki dependency ke layer lain
   - Pure business logic

2. **UseCase Layer** (`internal/usecase/`)
   - Berisi business logic aplikasi
   - Menggunakan domain models
   - Independent dari framework

3. **Interface Layer** (`internal/interfaces/`)
   - Berisi HTTP handlers dan routing
   - Menerjemahkan HTTP request ke usecase
   - Mengkonversi response ke DTO

4. **Infrastructure Layer** (`internal/infrastructure/`)
   - Berisi implementasi database, cache, external services
   - Repository pattern untuk data access
   - Isolated dari business logic

5. **Shared Layer** (`internal/shared/`)
   - Berisi utilities yang digunakan di semua layer
   - Config, error handling, validators
   - Common constants

## API Endpoints

```
GET    /                  # Welcome message
GET    /users             # Get all users
GET    /users/:id         # Get user by ID
POST   /users             # Create new user
PUT    /users/:id         # Update user
DELETE /users/:id         # Delete user
```

## Setup & Running

### Prerequisites
- Go 1.22+
- PostgreSQL 12+

### Configuration

Copy `resources/config.json.example` ke `resources/config.json`:

```bash
cp resources/config.json.example resources/config.json
```

Edit file `resources/config.json` dengan konfigurasi database Anda:

```json
{
  "database": {
    "host": "localhost",
    "port": "5432",
    "user": "postgres",
    "password": "your_password",
    "dbname": "mogi",
    "sslmode": "disable"
  },
  "server": {
    "port": "8000"
  }
}
```

### Install Dependencies

```bash
go mod download
go mod tidy
```

### Run Application

```bash
make run
```

atau:

```bash
go run main.go
```

Server akan berjalan di `localhost:8000`

## Request Examples

### Get All Users
```bash
curl http://localhost:8000/users
```

### Create User
```bash
curl -X POST http://localhost:8000/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword"
  }'
```

### Get User by ID
```bash
curl http://localhost:8000/users/1
```

### Update User
```bash
curl -X PUT http://localhost:8000/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newemail@example.com",
    "password": "newsecurepassword"
  }'
```

### Delete User
```bash
curl -X DELETE http://localhost:8000/users/1
```

## Best Practices

1. **Separation of Concerns** - Setiap layer memiliki tanggung jawab spesifik
2. **Dependency Injection** - Dependencies diinject, bukan diciptakan
3. **Interface-based Design** - Gunakan interfaces untuk loose coupling
4. **Error Handling** - Gunakan custom error types untuk better error handling
5. **DTO Pattern** - Pisahkan internal models dengan external DTOs

## Development Guidelines

### Menambah Feature Baru

1. Buat domain model di `internal/domain/`
2. Buat DTOs di `internal/dto/`
3. Buat repository di `internal/infrastructure/`
4. Buat usecase di `internal/usecase/`
5. Buat HTTP handler di `internal/interfaces/handler/`
6. Register routes di `internal/interfaces/server/server.go`

### Testing

Tests harus ditempatkan di sebelah file yang di-test dengan naming convention `*_test.go`

```bash
go test ./...
```

## References

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [SOLID Principles](https://en.wikipedia.org/wiki/SOLID)
- [Go Project Layout](https://github.com/golang-standards/project-layout)

