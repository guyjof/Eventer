[![made-for-macOS](https://img.shields.io/badge/Made%20for-macOS-1f425f.svg?logo=apple)](https://www.apple.com/macos/)
![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg?logo=go)
![made-with-Gin](https://img.shields.io/badge/Made%20with-Gin-1f425f.svg?logo=go)
[![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](http://unlicense.org/)

# Eventer API

Eventer is a RESTful API built with Go and Gin framework for managing events. It allows users to sign up, log in, create events, and register for events.


## Getting Started

### Prerequisites

- Go 1.16 or higher
- A running instance of a SQL database (e.g., MySQL, PostgreSQL)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/guyjof/Eventer.git
    cd eventer-api
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

### Running the Application

Start the server:

```sh
go run main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

### Health Check

- `GET /readiness-check`
- `GET /liveness-check`

### User Routes

- `POST /signup` - Sign up a new user
- `POST /login` - Log in an existing user

### Event Routes

- `GET /events` - Get all events
- `GET /events/:id` - Get a specific event by ID
- `POST /events` - Create a new event (requires authentication)
- `PUT /events/:id` - Update an event (requires authentication)
- `DELETE /events/:id` - Delete an event (requires authentication)
- `POST /events/:id/register` - Register for an event (requires authentication)
- `DELETE /events/:id/register` - Cancel registration for an event (requires authentication)

## Middleware

- `middlewares/auth.go` - Authentication middleware to protect routes

## Models

- `models/event.go` - Event model and database operations
- `models/user.go` - User model and database operations

## Utilities

- `utils/hash.go` - Utility functions for hashing
- `utils/jwt.go` - Utility functions for JWT token generation and validation
