# Niavo

> **Turn intention into execution.**

Niavo is an execution platform that helps organizations manage work from request to completion. It provides a simple and flexible way to organize tasks, assign responsibilities, and track progress across teams.

---

## Features (MVP)

- User Authentication
- Organizations
- Work Items
- Workflows
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

```
niavo/
├── api/
├── backend/
├── deployments/
├── docs/
├── frontend/
├── tests/
└── README.md
```

---

## Getting Started

### Prerequisites

Install:

- Go 1.22+
- Node.js 20+
- PostgreSQL
- Docker (optional)

---

### Clone the Repository

```bash
git clone git@github.com:breodoyo/niavo.git
cd niavo
```

---

### Backend

```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

The backend will be available at:

```
http://localhost:8080
```

---

### Frontend

```bash
cd frontend
npm install
npm run dev
```

The frontend will be available at:

```
http://localhost:5173
```

---

### Using Docker

```bash
docker compose up --build
```

---

## Documentation

Project documentation can be found in the `docs/` directory:

- `vision.md`
- `prd.md`
- `architecture.md`

---

## Development Status

🚧 MVP under active development.

Current focus:

- Authentication
- Organizations
- Work Items
- Workflows

---

## Roadmap

- Authentication
- Organizations
- Work Items
- Dashboard
- Workflow Management
- Automation
- AI Assistance
- Third-party Integrations

---

## Author

**Brender Odoyo**

- GitHub: https://github.com/breodoyo

---

## License

This project is licensed under the MIT License.