# Niavo

> **Turn intention into execution.**

Niavo is an execution platform that helps organizations manage work from request to completion. It enables teams to organize, assign, and track work through customizable workflows, replacing scattered spreadsheets, emails, and chat messages with a single source of truth.

---

## Features (MVP)

- User Authentication
- Organizations
- Work Items
- Workflow Management
- Dashboard
- REST API

---

## Tech Stack

### Backend
- Go
- PostgreSQL
- JWT Authentication

### Frontend
- React
- TypeScript

### Infrastructure
- Docker
- Docker Compose

---

## Project Structure

```text
niavo/
├── .github/         # GitHub Actions workflows
├── api/             # OpenAPI specification
├── backend/         # Go backend
├── deployments/     # Deployment configuration
├── docs/            # Project documentation
├── frontend/        # React frontend
├── tests/           # Integration and E2E tests
├── .env.example
├── docker-compose.yml
└── README.md
```

---

## Getting Started

### Prerequisites

- Go 1.22+
- Node.js 20+
- PostgreSQL
- Docker (optional)

### Clone the Repository

```bash
git clone git@github.com:breodoyo/niavo.git
cd niavo
```

### Backend

```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```

### Docker

```bash
docker compose up --build
```

---

## Documentation

| Document | Description |
|----------|-------------|
| `vision.md` | Product vision and mission |
| `prd.md` | Product requirements document |
| `architecture.md` | Technical architecture |

---

## Roadmap

- Authentication
- Organizations
- Work Items
- Dashboard
- Workflow Management
- Automation
- AI Assistance
- Integrations

---

## Author

**Brender Odoyo**

GitHub: https://github.com/breodoyo

---

## License

Licensed under the MIT License.