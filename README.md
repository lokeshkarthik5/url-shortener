# Go URL Shortener

URL shortener built in **Go**, designed with clean architecture principles

---

## Features

- Create short URLs
- Retrieve original URLs
- Update existing URLs
- Delete URLs
- URL statistics
- Health check endpoint
- Request metrics middleware
- PostgreSQL persistence

---

## Tech Stack

- **Go** (net/http)
- **PostgreSQL**
- **sqlc** for type-safe queries
- **UUIDs**
- Standard library middleware
- Clean, layered architecture

---

## How To Run

### Prerequisites

- Go 1.21+
- PostgreSQL
- `sqlc`
- `goose` (or any migration tool)

## Environment Variables

```env
DATABASE_URL=postgres://user:password@localhost:5432/url_shortener?sslmode=disable
PORT=3001
```

### Run migrations:

```bash 
goose -dir migrations postgres "$DATABASE_URL" up
```

### Generate sqlc code:

```bash
sqlc generate
```

### Running the Server:

```bash
go run cmd/api/main.go
```

The server will run at 3001
