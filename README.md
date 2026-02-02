# Go Workspace Projects

This workspace contains a collection of Go projects, each focusing on a specific concept or feature of the Go programming language. The projects progress from fundamentals to advanced topics including REST APIs, concurrency, and full-stack applications.

---

## Project Structure

### 01hello/
- **Description**: A simple "Hello, World!" program to get started with Go.
- **Files**: `main.go`

### 02variable/
- **Description**: Demonstrates variable declaration and usage in Go.
- **Files**: `main.go`

### 03userinput/
- **Description**: Handles user input and demonstrates how to read from the console.
- **Files**: `go.mod`, `main.go`

### 04conversion/
- **Description**: Covers type conversion in Go.
- **Files**: `main.go`

### 05mytime/
- **Description**: Explores working with time and dates in Go.
- **Files**: `go.mod`, `main.go`

### 06mypointers/
- **Description**: Introduces pointers and their usage in Go.
- **Files**: `main.go`

### 07array/
- **Description**: Demonstrates arrays and their operations in Go.
- **Files**: `go.mod`, `main.go`

### 08myslices/
- **Description**: Explains slices, a more flexible alternative to arrays in Go.
- **Files**: `main.go`

### 09mymaps/
- **Description**: Covers maps, a key-value data structure in Go.
- **Files**: `main.go`

### 10mystructs/
- **Description**: Introduces structs, a way to define custom data types in Go.
- **Files**: `main.go`

### 11ifelse/
- **Description**: Demonstrates conditional statements (`if-else`) in Go.
- **Files**: `main.go`

### 12switchcase/
- **Description**: Covers the `switch` statement for conditional logic in Go.
- **Files**: `main.go`

### 13loops/
- **Description**: Explains looping constructs like `for` in Go.
- **Files**: `main.go`

### 14functions/
- **Description**: Introduces functions and their usage in Go.
- **Files**: `main.go`

### 15methods/
- **Description**: Demonstrates methods and their association with structs in Go.
- **Files**: `main.go`

### 16defer/
- **Description**: Explains the `defer` statement and its use cases in Go.
- **Files**: `main.go`

### 17files/
- **Description**: Covers file handling, including reading and writing files in Go.
- **Files**: `main.go`, `myfile.txt`

### 18webrequests/
- **Description**: Demonstrates how to make HTTP web requests in Go.
- **Files**: `main.go`

### 19urls/
- **Description**: Explains working with URLs and parsing them in Go.
- **Files**: `main.go`

### 20tracker/
- **Description**: CLI expense tracker with JSON file persistence. Supports add, list, update, delete, and summary commands with optional month filtering.
- **Files**: `main.go`, `expenses.json`
- **Usage**: `go run main.go [add|update|delete|list|summary] [flags]`

### 21bitmorejson/
- **Description**: Advanced JSON encoding and decoding. Demonstrates struct tags (`json:"coursename"`, `json:"-"`, `omitempty`), Marshal/Unmarshal, and dynamic parsing with `map[string]interface{}`.
- **Files**: `main.go`

### 22mymod/
- **Description**: Introduces Go modules, `go mod` commands, and vendoring. Uses Gorilla Mux for basic HTTP routing with a simple home handler.
- **Files**: `main.go`, `go.mod`, `go.sum`, `vendor/`

### 23api/
- **Description**: REST API for a course catalog. Full CRUD operations using Gorilla Mux with in-memory storage.
- **Endpoints**: `GET /`, `GET /courses`, `GET /course/{id}`, `POST /course`, `PUT /course/{id}`, `DELETE /course/{id}`
- **Files**: `main.go`, `go.sum`
- **Port**: 4000

### 24context/
- **Description**: Demonstrates the `context` package for cancellation and timeouts. Uses `context.WithTimeout` to stop a goroutine after a specified duration.
- **Files**: `main.go`

### 25mongoapi/
- **Description**: MongoDB-backed API for a Netflix watchlist. CRUD operations for movies with MongoDB Atlas connection.
- **Endpoints**: `GET /api/movies`, `POST /api/movie`, `PUT /api/movie/{id}`, `DELETE /api/movie/{id}`, `DELETE /api/deleteallmovie`
- **Files**:
  - `main.go` – Entry point
  - `controller/control.go` – HTTP handlers and MongoDB operations
  - `model/models.go` – Netflix struct (ID, Movie, Watched)
  - `router/routers.go` – Route definitions
  - `go.mod`, `go.sum`
- **Port**: 4000

### 26goroutines/
- **Description**: Concurrent HTTP status checker using goroutines. Fetches status codes from multiple websites with `sync.WaitGroup` and `sync.Mutex` for thread-safe access.
- **Files**: `main.go`, `go.mod`

### 27mutexANDawaitgrp/
- **Description**: Demonstrates race conditions and synchronization. Uses `sync.RWMutex` (Lock/RLock) and `sync.WaitGroup` for safe concurrent reads and writes.
- **Files**: `main.go`, `go.mod`

### 28channels/
- **Description**: Introduces Go channels for goroutine communication. Covers buffered channels, send/receive-only channels, and closing channels.
- **Files**: `main.go`, `go.mod`, `go.sum`

### todo-api/
- **Description**: Full REST API for todos with JWT authentication. Uses Gin, GORM, and SQLite.
- **Files**:
  - `main.go` – Entry point
  - `config/config.go` – Environment configuration
  - `database/database.go` – GORM/SQLite connection
  - `controllers/auth_controller.go` – Register, Login
  - `controllers/todo_controller.go` – CRUD for todos
  - `middleware/auth.go` – JWT authentication middleware
  - `models/todo.go` – Todo model
  - `models/user.go` – User model
  - `routes/routes.go` – Route registration
  - `utils/jwt.go` – JWT utilities
  - `go.mod`, `go.sum`, `todo.db`
- **Endpoints**: `POST /register`, `POST /login`, `POST /todos`, `GET /todos`, `PUT /todos/:id`, `DELETE /todos/:id` (protected)
- **Port**: 8080

---

## How to Run

1. Navigate to the desired project directory.
2. For projects with `go.mod`, run:
   ```sh
   go mod tidy   # If dependencies need to be fetched
   go run main.go
   ```
3. For projects without `go.mod`, run from the workspace root or ensure the module path is correct:
   ```sh
   go run main.go
   ```

---

## Quick Reference

| Project      | Topic                     | Key Concepts                        |
|-------------|---------------------------|-------------------------------------|
| 01-04       | Basics                    | Hello, variables, input, conversion |
| 05-10       | Data structures           | Time, pointers, arrays, slices, maps, structs |
| 11-16       | Control flow & functions  | if/else, switch, loops, functions, methods, defer |
| 17-19       | I/O & networking          | Files, HTTP requests, URLs          |
| 20-21       | JSON & CLI                | JSON encoding/decoding, CLI apps    |
| 22-23       | Modules & APIs            | Go modules, REST APIs               |
| 24          | Context                   | Cancellation, timeouts              |
| 25          | MongoDB                   | NoSQL, MongoDB driver               |
| 26-28       | Concurrency               | Goroutines, mutex, channels         |
| todo-api    | Full application          | Gin, GORM, JWT, SQLite              |
