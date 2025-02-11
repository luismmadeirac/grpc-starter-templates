# Microservice Templates

This repository contains a collection of microservice templates for different programming languages and frameworks. Each template follows best practices and includes common patterns for building production-ready microservices.

## Available Templates

- **Go** - Standard Go microservice with gRPC
- **Node.js** - Express-based microservice
- **NestJS** - TypeScript-based microservice using NestJS framework
- **Python** - Python microservice using gRPC
- **C++** - C++ microservice with gRPC support
- **Java** - Spring Boot-based microservice

## Features

All templates include:

- **Health Checks**: Both HTTP and gRPC health check implementations
- **Metrics**: Prometheus integration for monitoring
- **Database**: Configured database connections with connection pooling
- **Configuration**: Environment-based configuration management
- **Logging**: Structured logging setup
- **Middleware**: CORS, Rate limiting, and other common middleware
- **Docker**: Containerization support
- **Kubernetes**: Production-ready K8s configurations
- **Testing**: Unit and integration test setups

## Getting Started

1. Choose a template that matches your technology stack
2. Use the service generation script:

   ```bash
   make new-service SERVICE_TYPE=<language> SERVICE_NAME=<name>
   ```

3. Follow the language-specific README in the generated service directory

## Template Structure

Each template follows a standard structure:

```
service-name/
├── cmd/ or src/         # Main application entry point
├── internal/ or lib/    # Internal packages
│   ├── config/         # Configuration management
│   ├── database/       # Database connections and models
│   ├── grpc/          # gRPC server implementation
│   ├── metrics/       # Metrics and monitoring
│   └── middleware/    # Common middleware
├── proto/             # Protocol buffer definitions
└── Dockerfile        # Container definition
```

## Infrastructure

Templates include infrastructure configurations:

- Kubernetes manifests
- Docker configurations
- Database setups
- Monitoring configurations (Prometheus/Grafana)

## Development

### Prerequisites

- Docker
- Kubernetes (local cluster, e.g., minikube)
- Protocol Buffers compiler
- Make

### Local Development

1. Clone the template repository
2. Generate a new service:

   ```bash
   make new-service SERVICE_TYPE=go SERVICE_NAME=my-service
   ```

3. Navigate to the service directory
4. Follow language-specific development instructions

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

### Adding New Templates

1. Follow the existing template structure
2. Include all required components (health checks, metrics, etc.)
3. Add appropriate documentation
4. Update the service generation script

## License

This project is licensed under the MIT License - see the LICENSE file for details.
