# Go Web Server With Real Time

This project is a web server built with Go, using the Gin framework for handling HTTP requests, JWT for authentication, and WebSockets for real-time communication. It includes user registration and login functionality, and a secure WebSocket connection.

## Features

- User registration and login
- JWT-based authentication
- WebSocket for real-time communication
- PostgreSQL database integration with GORM
- Middleware for JWT authentication

## Getting Started

### Prerequisites

- Go (version 1.22 or later)
- PostgreSQL

### Installation

1.  Clone the repository:

    ```sh
    git clone https://github.com/Jasveer399/Chat-To.git
    cd Chat-To
    ```

2.  Install dependencies:

    ```sh
    go mod tidy
    ```

3.  Set up the database:
    - Create a `.env` file in the root directory.
    - Add the following environment variables to the `.env` file:
      ```
      DB_HOST=localhost
      DB_PORT=5432
      DB_USER=your_db_user
      DB_PASSWORD=your_db_password
      DB_NAME=your_db_name
      ```

### Running the Application

To run the application, execute the following command:

```sh
go run main.go
```

The server will start on `http://localhost:3000`.

## API Endpoints

### HTTP Endpoints

- `POST /register`: Register a new user.
- `POST /login`: Log in an existing user.
- `GET /get-user`: Get all users (requires JWT authentication).

### WebSocket Endpoint

- `GET /ws`: Establish a WebSocket connection (requires JWT authentication).

## Project Structure

```
.
├── common
│   └── claims.go
├── controllers
│   └── auth
│       └── handlers.go
├── database
│   ├── db.go
│   └── migrate.go
├── go.mod
├── go.sum
├── main.go
├── middleware
│   ├── getUserID.go
│   └── jwtAuth.go
├── models
│   └── message.go
├── utils
│   └── response.go
└── websocket
    ├── client.go
    ├── handler.go
    └── hub.go
```

## Dependencies

- [GORM](https://gorm.io/): ORM library for Go.
- [PostgreSQL Driver for GORM](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL): PostgreSQL database driver for GORM.
- [JWT for Go](https://github.com/golang-jwt/jwt): JSON Web Token implementation for Go.
- [Gorilla WebSocket](https://github.com/gorilla/websocket): WebSocket implementation for Go.
