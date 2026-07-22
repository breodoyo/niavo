# Architecture

## Overview

Niavo follows a client-server architecture.

The frontend communicates with the backend through a REST API. The backend contains the business logic and interacts with a PostgreSQL database for persistent storage.

```text
┌───────────────┐
│   Frontend    │
│ React + TS    │
└───────┬───────┘
        │ HTTP/REST
        ▼
┌───────────────┐
│    Backend    │
│      Go       │
└───────┬───────┘
        │
        ▼
┌───────────────┐
│ PostgreSQL DB │
└───────────────┘
```

---

## Technology Stack

### Frontend

- React
- TypeScript

### Backend

- Go
- REST API
- JWT Authentication

### Database

- PostgreSQL

### Infrastructure

- Docker
- Docker Compose

---

## Repository Structure

```text
niavo/
├── api/
├── backend/
├── deployments/
├── docs/
├── frontend/
└── tests/
```

---

## Backend Structure

```text
backend/
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── auth/
│   ├── common/
│   ├── config/
│   ├── database/
│   ├── middleware/
│   ├── organization/
│   ├── user/
│   ├── workflow/
│   └── workitem/
│
├── migrations/
├── pkg/
├── scripts/
├── go.mod
└── go.sum
```

---

## Directory Responsibilities

| Directory | Responsibility |
|-----------|----------------|
| `cmd/` | Application entry point |
| `internal/` | Business logic |
| `auth/` | Authentication and authorization |
| `organization/` | Organization management |
| `user/` | User management |
| `workitem/` | Work item lifecycle |
| `workflow/` | Workflow definitions and transitions |
| `database/` | Database connection and repositories |
| `middleware/` | HTTP middleware |
| `common/` | Shared utilities |
| `migrations/` | Database schema migrations |
| `pkg/` | Reusable packages |
| `scripts/` | Development and maintenance scripts |

---

## Design Principles

- Domain-driven project organization
- Separation of concerns
- Modular architecture
- RESTful API design
- Secure authentication
- Scalable and maintainable codebase