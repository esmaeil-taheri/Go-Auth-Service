# Authentication API

A simple authentication project built with Go, Echo, and MySQL. It supports user registration, login with JWT access and refresh tokens, and an authenticated profile endpoint.

## Run locally

Install Go (see `go.mod` for the required version) and Docker with Docker Compose.

```sh
docker compose up -d
go run .
```

Wait for MySQL to become healthy before starting the application. Database migrations run automatically on startup. The API listens on `http://localhost:8080`.

The application currently uses the settings defined in `main.go`, including the database connection and JWT signing key. Although it loads `config.yml`, those loaded settings are not yet used to configure the services. Replace the development credentials and signing key before deploying.

## Endpoints

| Method | Path | Description |
| --- | --- | --- |
| GET | `/health-check` | Check service health |
| POST | `/users/register` | Register with `name`, `phone_number`, and `password` |
| POST | `/users/login` | Log in with `phone_number` and `password` |
| GET | `/users/profile` | Get the current user's profile |

Send JSON request bodies with `Content-Type: application/json`. For the profile endpoint, include the access token in the `Authorization: Bearer <access_token>` header.

Login returns both access and refresh tokens. There is currently no endpoint for exchanging a refresh token.

## Project architecture

This is a single Go application with a layered architecture. HTTP handling, authentication logic, and database access live in separate packages, while `main.go` creates their dependencies and connects them together.

| Package | Responsibility |
| --- | --- |
| `delivery/httpserver` | Start Echo, register routes, and install middleware |
| `delivery/httpserver/userhandler` | Decode requests, call validation and services, and return HTTP responses |
| `delivery/httpserver/middleware` | Authenticate requests to protected routes |
| `validator/uservalidator` | Validate request fields and check phone number uniqueness |
| `service/userservice` | Implement registration, login, and profile retrieval |
| `service/authservice` | Create and parse JWTs |
| `repository/mysql` | Execute SQL queries and map rows to users |
| `repository/migrator` | Apply database schema migrations |
| `entity` | Define the user model |
| `param` | Define request and response structures |
| `pkg` | Share password hashing and error handling helpers |
| `config` | Define configuration types and load settings |

The normal request flow is:

```text
HTTP request
  -> Echo route and middleware
  -> handler -> request validator
  -> user service -> repository -> MySQL
  -> handler -> JSON response
```

Registration validates the submitted fields, checks whether the phone number is already registered, hashes the password with bcrypt, and saves the user. Login loads the user, compares the password hash, and creates access and refresh tokens. Profile requests pass through JWT middleware before the handler requests the user identified by the token.

The user service declares small `Repository` and `AuthGenerator` interfaces for the operations it needs. The MySQL repository and authentication service implement those interfaces. This keeps SQL and Echo out of the user service and allows service tests to substitute small fake implementations. Dependencies are passed through constructors; a dependency injection framework is unnecessary here.
