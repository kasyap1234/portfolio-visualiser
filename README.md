# Portfolio Visualizer

A comprehensive portfolio visualiser that tracks stocks, bank balances, and provides detailed analysis.

## Tech Stack

- **Frontend:** SvelteKit (TypeScript)
- **Backend:** Go
- **Database:** PostgreSQL
- **Cache:** Redis
- **Automation:** Taskfile

## Getting Started

### Prerequisites

- Go (1.21+)
- Node.js (v18+)
- Docker & Docker Compose
- Task (optional, but recommended: `go install github.com/go-task/task/v3/cmd/task@latest`)

### Initial Setup

1. Initialize the project dependencies:
   ```bash
   task init
   ```

2. Start the infrastructure (PostgreSQL, Redis):
   ```bash
   task infra:up
   ```

### Development

- **Run Backend:**
  ```bash
  task backend:dev
  ```

- **Run Frontend:**
  ```bash
  task frontend:dev
  ```

## Folder Structure

- `backend/`: Go source code
  - `cmd/`: Entry points
  - `internal/`: Business logic, providers, repositories
  - `pkg/`: Public library code
- `client/`: SvelteKit frontend
- `deployments/`: Infrastructure configurations (Docker Compose)
- `Taskfile.yml`: Automation tasks
