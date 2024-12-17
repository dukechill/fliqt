# FLIQT

FLIQT is a modular Human Resource Management System backend service built with **Go (Golang)**, using the **Gin** framework to implement RESTful APIs. It integrates **MySQL** as the database, **Redis** as the caching system, and **MinIO** for object storage functionality.

This project supports containerized deployment using **Docker Compose** to quickly start all required services. It also features database migrations and unit tests.

---

## **Table of Contents**

1. [Project Features](#project-features)
2. [Technology Stack](#technology-stack)
3. [System Requirements](#system-requirements)
4. [Installation and Execution](#installation-and-execution)
5. [Main Commands](#main-commands)
6. [Environment Variables](#environment-variables)
7. [Project Structure](#project-structure)
8. [Future Enhancements](#future-enhancements)

---

## **Project Features**

- **RESTful API**: Built with the high-performance Gin framework.
- **Database**: MySQL integration with GORM for database operations and migrations.
- **Caching System**: Redis for improved system performance.
- **Object Storage**: MinIO provides S3-compatible storage.
- **Containerized Deployment**: Quickly launch integrated services using Docker and Docker Compose.
- **Version Management**: Automatic version tagging for easy maintenance.
- **Unit Testing**: Built-in testing examples to ensure system stability.

---

## **Technology Stack**

- **Language**: Go (Golang)
- **Framework**: Gin
- **Database**: MySQL + GORM
- **Cache**: Redis
- **Object Storage**: MinIO
- **Containerization**: Docker, Docker Compose
- **Testing**: Go Testing Framework
- **Version Control**: Git

---

## **System Requirements**

- **Go** 1.20 or higher
- **Docker** 20.10 or higher
- **Docker Compose** 1.29 or higher
- **Git**

---

## **Installation and Execution**

### **1. Clone the Project**

```bash
git clone https://github.com/your-username/fliqt.git
cd fliqt
```

### **2. Build and Run**

Use Docker Compose to launch all services:

```bash
make docker-compose-run
```

This will start:
- MySQL
- Redis
- MinIO
- FLIQT API service

The API will run on `http://localhost:8080`.

---

## **Main Commands**

| Command                   | Description                         |
|---------------------------|-------------------------------------|
| `make dev-db-migrate`     | Run database migrations             |
| `make dev-db-rollback`    | Roll back the last database migration |
| `make docker-build`       | Build and push Docker image         |
| `make docker-compose-run` | Start all services with Docker Compose |
| `make clean`              | Clean up Docker resources           |
| `make version`            | Display project version information |

---

## **Environment Variables**

Create a `.env` file in the project root directory and fill it with the following variables:

```env
DB_HOST=mysql
DB_PORT=3306
DB_USER=root
DB_PASSWORD=rootpassword
DB_NAME=fliqt

REDIS_URL=redis://redis:6379
S3_ENDPOINT=http://minio:9000
S3_BUCKET=fliqt
S3_KEY=minioadmin
S3_SECRET=minioadmin
DEBUG=true
```

---

## **Project Structure**

```plaintext
fliqt/
│
├── Dockerfile                 # Docker build configuration
├── docker-compose.yml         # Multi-service orchestration
├── makefile                   # Convenient command execution
├── main.go                    # Project entry point
├── go.mod                     # Go module configuration
├── config/                    # Configuration logic
├── internal/                  
│   ├── handler/               # API request handlers
│   └── repository/            # Data access logic
└── tests/                     # Test code
```

---

## **Future Enhancements**

- Add user authentication and authorization (JWT).
- Provide more unit tests to improve test coverage.
- Integrate CI/CD pipelines for automated build and deployment.
- Enhance API documentation with tools like Swagger.
- Optimize Docker build for production environments.

---

## **Contact Information**

For any questions or suggestions, please contact [your-email@example.com].

---

**Project Maintainer**: `your-name`  
**Version**: `v1.0.0`  
**License**: MIT
