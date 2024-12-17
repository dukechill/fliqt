# FLIQT

FLIQT is a backend service for a Human Resource Management System, created as a demonstration project for interview purposes. The project includes basic functionalities and can be easily started using containerized deployment.

---

## **Table of Contents**

1. [Project Features](#project-features)
2. [System Requirements](#system-requirements)
3. [Installation and Execution](#installation-and-execution)
4. [Main Commands](#main-commands)
5. [Environment Variables](#environment-variables)
6. [Project Structure](#project-structure)

---

## **Project Features**

- Basic functionalities for a Human Resource Management System.
- Containerized deployment for easy setup.
- Supports database migrations and basic testing.

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

The API will run on `http://localhost:8080`.

---

## **Main Commands**

| Command                   | Description                         |
|---------------------------|-------------------------------------|
| `make dev-db-migrate`     | Run database migrations             |
| `make dev-db-rollback`    | Roll back the last database migration |
| `make docker-compose-run` | Start all services with Docker Compose |
| `make clean`              | Clean up Docker resources           |

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
- Enhance API documentation with tools like Swagger.

---

**Project Maintainer**: `your-name`  
**Version**: `v1.0.0`  
**License**: MIT
