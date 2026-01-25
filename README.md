# Student API

A simple RESTful API for managing student records, built with Go (Golang) and SQLite. This project demonstrates basic CRUD operations, configuration management, and structured logging in Go.

## Features

-   **Student Management**: Create and retrieve student records.
-   **SQLite Database**: Lightweight and embedded data storage.
-   **Configuration**: Easy configuration via YAML files or environment variables.
-   **Validation**: Request data validation using `go-playground/validator`.
-   **Structured Logging**: Uses `log/slog` for structured logging.

## Tech Stack

-   **Language**: Go 1.24
-   **Database**: SQLite
-   **Router**: `net/http` (Standard Library `ServeMux`)
-   **Config**: `ilyakaznacheev/cleanenv`
-   **Validation**: `go-playground/validator/v10`

## Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/CodeMaverick-143/Golang_learning.git
    cd Golang_learning
    ```

2.  **Download dependencies:**
    ```bash
    go mod download
    ```

## Configuration

The application is configured using a YAML file located at `config/local.yaml` by default. You can also configure it using environment variables.

**Default Configuration (`config/local.yaml`):**
```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server: 
  address: "localhost:8082"
```

**Environment Variables:**
-   `CONFIG_PATH`: Path to the config file (default: `config/local.yaml`).
-   `ENV`: Environment (e.g., `local`, `dev`, `prod`).
-   `STORAGE_PATH`: Path to the SQLite database file.
-   `HTTP_SERVER_ADDRESS`: Address for the HTTP server to listen on.

## Usage

### Run the Server

To start the server using the default configuration:

```bash
go run cmd/student-api/main.go
```

You should see output indicating the server has started:
```text
level=INFO msg="Database setup successfully" ...
level=INFO msg="Server started" ...
```

## API Endpoints

### 1. Create Student

-   **URL**: `/api/students`
-   **Method**: `POST`
-   **Content-Type**: `application/json`

**Request Body:**

```json
{
  "name": "Arpit Sarang",
  "email": "arpitsarang@gmail.com",
  "age": 20
}
```

**Response:**

```json
{
  "id": "1"
}
```

### 2. Get Student by ID

-   **URL**: `/api/students/{id}`
-   **Method**: `GET`

**Response:**

```json
{
  "id": 1,
  "name": "Arpit Sarang",
  "email": "arpitsarang@gmail.com",
  "age": 20
}
```

## Database

The application uses SQLite. The database file will be automatically created at `storage/storage.db` (or the path specified in your config) when the application starts.