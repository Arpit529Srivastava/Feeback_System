# FeedBack System

A full-stack feedback management system with a Go backend and a Next.js (TypeScript) frontend.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Docker Usage](#docker-usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- Collect and manage user feedback
- RESTful API backend (Go)
- Modern frontend (Next.js + TypeScript)
- Dockerized for easy deployment

## Tech Stack

- **Backend:** Go, Gin (or similar), Docker
- **Frontend:** Next.js, React, TypeScript, Docker

## Getting Started

### Backend Setup

1. **Install Go** (if not already installed):  
   https://golang.org/doc/install

2. **Install dependencies:**
   ```bash
   cd backend
   go mod tidy
   ```

3. **Run the backend server:**
   ```bash
   go run main.go
   ```

   The backend should now be running (default: `localhost:8080`).

### Frontend Setup

1. **Install Node.js** (if not already installed):  
   https://nodejs.org/

2. **Install dependencies:**
   ```bash
   cd frontend
   npm install
   ```

3. **Run the frontend development server:**
   ```bash
   npm run dev
   ```

   The frontend should now be running (default: `localhost:3000`).

## Docker Usage

You can run both backend and frontend using Docker.

### Backend

```bash
cd backend
docker build -t feedback-backend .
docker run -p 8080:8080 feedback-backend
```

### Frontend

```bash
cd frontend
docker build -t feedback-frontend .
docker run -p 3000:3000 feedback-frontend
```

## Project Structure

```
FeedBack_System/
├── backend/
│   ├── main.go
│   ├── controllers/
│   ├── models/
│   ├── routes/
│   ├── config/
│   ├── go.mod
│   └── Dockerfile
└── frontend/
    ├── app/
    ├── components/
    ├── public/
    ├── package.json
    ├── next.config.ts
    └── Dockerfile
```

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](LICENSE) (or your preferred license)
