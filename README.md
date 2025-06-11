# Feedback System

A production-ready, full-stack feedback management system with Go backend, MongoDB database, and Next.js frontend. Designed for scalable cloud deployment with comprehensive DevOps practices.

## 🚀 Quick Start

```bash
# Clone and start with Docker Compose
git clone <repo-url>
cd FeedBack_System
docker-compose up -d

# Access services
# Frontend: http://localhost:3000
# Backend API: http://localhost:8080
# MongoDB: localhost:27017
```

## 🏗️ Architecture

- **Backend**: Go 1.21+ with Gin framework
- **Database**: MongoDB with persistent volumes
- **Frontend**: Next.js 14 with TypeScript
- **Containerization**: Docker & Docker Compose
- **Infrastructure**: AWS ECS/EKS (see [Infrastructure Repository](https://github.com/arpit529srivastava/Feedback_system_Infra))

## 📁 Project Structure

```
FeedBack_System/
├── backend/                 # Go API service
│   ├── cmd/                # Application entrypoints
│   ├── internal/           # Private application code
│   ├── pkg/               # Public libraries
│   ├── configs/           # Configuration files
│   ├── deployments/       # Docker & K8s configs
│   └── Dockerfile
├── frontend/              # Next.js application
│   ├── src/
│   ├── public/
│   ├── __tests__/         # Jest tests
│   └── Dockerfile
├── docker-compose.yml     # Local development stack
├── .github/workflows/     # CI/CD pipelines
└── docs/                  # API documentation
```

## 🔧 Development

### Prerequisites
```bash
# Required tools
go version          # >= 1.21
node --version      # >= 18
docker --version    # Latest
kubectl version     # For K8s deployments
```

### Local Development
```bash
# Backend
cd backend
go mod tidy
go run cmd/main.go

# Frontend  
cd frontend
npm install
npm run dev

# Full stack with Docker
docker-compose up --build
```

### Environment Configuration
```bash
# Backend (.env)
MONGODB_URI=mongodb://localhost:27017/feedback
PORT=8080
GIN_MODE=release

# Frontend (.env.local)
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## 🚀 Deployment

### Container Registry
```bash
# Build and push images
docker build -t feedback-backend ./backend
docker build -t feedback-frontend ./frontend

# Tag for ECR
docker tag feedback-backend:latest <account>.dkr.ecr.<region>.amazonaws.com/feedback-backend:latest
docker push <account>.dkr.ecr.<region>.amazonaws.com/feedback-backend:latest
```

### Infrastructure Deployment
Infrastructure is managed separately in the [Infrastructure Repository](https://github.com/arpit529srivastava/Feedback_system_Infra):

- **Terraform**: AWS ECS/VPC provisioning
- **Kubernetes**: Container orchestration
- **ECR**: Image registry
- **ECS Services**: Auto-scaling and load balancing

## 🧪 Testing & Quality

```bash
# Backend tests
cd backend
go test ./...
go test -race ./...
go test -cover ./...

# Frontend tests
cd frontend
npm test
npm run test:coverage
npm run lint
npm run type-check

# Integration tests
docker-compose -f docker-compose.test.yml up --abort-on-container-exit
```

## 📊 Monitoring & Observability

- **Health Checks**: `/health` endpoint for both services
- **Metrics**: Prometheus-compatible metrics
- **Logging**: Structured JSON logging
- **Tracing**: OpenTelemetry integration ready

## 🔒 Security

- **Authentication**: JWT-based auth
- **CORS**: Configured for production domains
- **Rate Limiting**: Built-in API rate limiting
- **Container Security**: Non-root containers, minimal base images
- **Infrastructure**: VPC isolation, security groups, IAM roles

## 🔄 CI/CD Pipeline

GitHub Actions workflows included:
- **Code Quality**: Linting, testing, security scanning
- **Build**: Multi-stage Docker builds
- **Deploy**: Automated ECS/EKS deployments
- **Rollback**: Blue-green deployment strategy

## 📈 Performance

- **Backend**: ~1000 RPS with proper tuning
- **Database**: Indexed queries, connection pooling
- **Frontend**: Static generation, code splitting
- **Caching**: Redis integration ready

## 🛠️ Configuration Management

```bash
# Environment-specific configs
configs/
├── development.yaml
├── staging.yaml
└── production.yaml

# Kubernetes configs in infrastructure repo
k8s/
├── backend/
└── frontend/
```

## 📚 API Documentation

- **OpenAPI/Swagger**: Available at `/swagger` endpoint
- **Postman Collection**: `docs/postman_collection.json`
- **API Docs**: Auto-generated from Go code comments

## 🤝 Contributing

1. Fork the repository
2. Create feature branch: `git checkout -b feature/new-feature`
3. Follow Go and TypeScript best practices
4. Add tests for new functionality
5. Update documentation
6. Submit pull request

## 📝 License

MIT License - see [LICENSE](LICENSE) file for details.

---

**Infrastructure Repository**: [Feedback System Infrastructure](https://github.com/arpit529srivastava/Feedback_system_Infra)

For infrastructure setup, deployment configurations, and cloud resource management.

### API Endpoints

-   `GET /api/feedback`: Get all feedback submissions.
-   `POST /api/feedback`: Create a new feedback submission.
-   `GET /health`: Check the health of the backend service.