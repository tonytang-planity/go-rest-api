# Go REST API

This project implements a RESTful API using Go, Gin, GORM, and PostgreSQL. It provides a basic setup for a users service where clients can create, read, update, and delete user records.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Ensure you have the following installed on your local machine:

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/tonytangdev/go-rest-api.git
cd go-rest-api
```

2. Start the services using Docker Compose

```bash
docker-compose up -d
```

### Usage

Start the server :

```bash
go run cmd/server/main.go
```

Once the services are up, the API will be available at http://localhost:8080. Here are some of the endpoints you can access:

    GET /users: List all users
    POST /users: Create a new user
    GET /users/:id: Get a user by ID
    PUT /users/:id: Update a user by ID
    DELETE /users/:id: Delete a user by ID

If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

    Go
    Gin
    GORM
    PostgreSQL
    Docker
