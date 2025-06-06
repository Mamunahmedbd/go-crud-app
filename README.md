# Go CRUD Application

A simple CRUD (Create, Read, Update, Delete) application built with Go, using Gin framework and GORM for database operations.

## Features

- RESTful API endpoints for User management
- CRUD operations
- Input validation
- JSON responses

## Prerequisites

- Go 1.16 or higher
- MySQL/PostgreSQL (depending on your database configuration)

## Installation

1. Clone the repository

```bash
git clone <your-repository-url>
cd go-crud-app
```

2. Install dependencies

```bash
go mod download
```

3. Set up your database configuration in `config/database.go`

4. Run the application

```bash
go run main.go
```

## API Endpoints

| Method | Endpoint   | Description       |
| ------ | ---------- | ----------------- |
| POST   | /users     | Create a new user |
| GET    | /users     | Get all users     |
| GET    | /users/:id | Get user by ID    |
| PUT    | /users/:id | Update user       |
| DELETE | /users/:id | Delete user       |

## Project Structure
